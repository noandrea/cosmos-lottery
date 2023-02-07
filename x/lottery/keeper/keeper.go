package keeper

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/crypto/blake2b"

	"lottery/x/lottery"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		bank       lottery.BankKeeper
	}
)

var (
	betsCounterKey        = []byte{0x1}
	maxLotteryBetKey      = []byte{0x2}
	minLotteryBetKey      = []byte{0x3}
	lotteryHashKey        = []byte{0x4}
	lotteryFeesBalanceKey = []byte{0x5}
	lotteryBetsBalanceKey = []byte{0x6}
	betKeyPrefix          = []byte{0x12}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bank lottery.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(lottery.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bank:       bank,
	}
}

type Payout int

const (
	PayoutFull Payout = iota
	PayoutBet
)

// Logger prepare the logger for the lottery module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", lottery.ModuleName))
}

// GetBalance retrieve a spendable balance for an account
func (k Keeper) GetBalance(ctx sdk.Context, account sdk.AccAddress, demon string) sdk.Int {
	return k.bank.SpendableCoins(ctx, account).AmountOf(demon)
}

// GetPoolSize retrieve the number of bets for the current lottery
func (k Keeper) GetPoolSize(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	return MustGetInt(store, betsCounterKey, sdk.ZeroInt()).Uint64()
}

// GetLotteryData retrieve the current lottery data necessary to calculate the winner and the payout
func (k Keeper) GetLotteryData(ctx sdk.Context) (lotteryHash []byte, maxBet, minBet sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	minBet = MustGetInt(store, minLotteryBetKey, sdk.ZeroInt())
	maxBet = MustGetInt(store, maxLotteryBetKey, sdk.ZeroInt())
	lotteryHash = store.Get(lotteryHashKey)
	return
}

// GetLotteryWinner retrieve the account associated with a bet index
func (k Keeper) GetLotteryWinner(ctx sdk.Context, winnerIndex uint64) (winnerAddress sdk.AccAddress, winnerBet lottery.Bet, err error) {
	store := ctx.KVStore(k.storeKey)

	bI := prefix.NewStore(store, betKeyPrefix).Iterator(nil, nil)
	defer bI.Close()

	for ; bI.Valid(); bI.Next() {
		// unmarshal the bet
		if err = json.Unmarshal(bI.Value(), &winnerBet); err != nil {
			err = fmt.Errorf("system error - failed to deserialize bet - %s", err.Error())
			return
		}
		// if it is not the winner, then continue
		if winnerBet.Index != winnerIndex {
			continue
		}
		// parse the account address
		addr := string(bI.Key())
		winnerAddress, err = sdk.AccAddressFromBech32(addr)
		if err != nil {
			err = fmt.Errorf("system error - failed to deserialize account address - %s", err.Error())
		}
		return
	}
	return
}

// ExecutePayout pays the amount to the winner and reset the lottery status
func (k Keeper) ExecutePayout(ctx sdk.Context, recipient sdk.AccAddress, payoutType Payout) error {
	store := ctx.KVStore(k.storeKey)
	var payoutAmount sdk.Int

	// determine the payout type
	switch payoutType {
	case PayoutFull:
		// pay everything and reset the balances
		payoutAmount = MustGetInt(store, lotteryBetsBalanceKey, sdk.ZeroInt()).Add(MustGetInt(store, lotteryFeesBalanceKey, sdk.ZeroInt()))
		MustStoreInt(store, lotteryBetsBalanceKey, sdk.ZeroInt())
		MustStoreInt(store, lotteryFeesBalanceKey, sdk.ZeroInt())
	case PayoutBet:
		// pay the bets and reset the bets balance
		payoutAmount = MustGetInt(store, lotteryBetsBalanceKey, sdk.ZeroInt())
		MustStoreInt(store, lotteryBetsBalanceKey, sdk.ZeroInt())
	}

	transferCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, payoutAmount))
	if err := k.bank.SendCoinsFromModuleToAccount(ctx, lottery.ModuleName, recipient, transferCoins); err != nil {
		return err
	}
	// reset the bets counter
	MustStoreInt(store, betsCounterKey, sdk.ZeroInt())
	// reset the bet list
	bI := prefix.NewStore(store, betKeyPrefix).Iterator(nil, nil)
	defer bI.Close()
	for ; bI.Valid(); bI.Next() {
		store.Delete(bI.Key())
	}
	// reset min/max bet value
	MustStoreInt(store, maxLotteryBetKey, sdk.NewIntFromUint64(math.MaxUint64))
	MustStoreInt(store, minLotteryBetKey, sdk.ZeroInt())
	// reset the lottery hash
	store.Set(lotteryHashKey, make([]byte, 256))

	return nil
}

// HasBet verify if an account has a bet placed for the current lottery
func (k Keeper) HasBet(ctx sdk.Context, account string) bool {
	store := ctx.KVStore(k.storeKey)
	key := append(betKeyPrefix, []byte(account)...)
	return store.Get(key) != nil
}

// PlaceBet place a bet for the current lottery
func (k Keeper) PlaceBet(ctx sdk.Context, account sdk.AccAddress, bet lottery.Bet, txBytes []byte) error {
	store := ctx.KVStore(k.storeKey)
	key := append(betKeyPrefix, []byte(account.String())...)
	if store.Get(key) != nil {
		return fmt.Errorf("bet already placed for account %s", account)
	}
	// transfer funds to the module account
	transferCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, bet.Amount.Add(bet.Fee)))
	if err := k.bank.SendCoinsFromAccountToModule(ctx, account, lottery.ModuleName, transferCoins); err != nil {
		return err
	}
	// store the transaction data into the store
	betBin, err := json.Marshal(bet)
	if err != nil {
		return fmt.Errorf("system error - failed to serialize bet - %s", err.Error())
	}
	store.Set(key, betBin)
	// increase the pool size
	MustStoreInt(store, betsCounterKey, sdk.NewIntFromUint64(bet.Index+1))
	// increase the bet balance
	balance := MustGetInt(store, lotteryBetsBalanceKey, sdk.ZeroInt()).Add(bet.Amount)
	MustStoreInt(store, lotteryBetsBalanceKey, balance)
	// increase the fee balance
	balance = MustGetInt(store, lotteryFeesBalanceKey, sdk.ZeroInt()).Add(bet.Fee)
	MustStoreInt(store, lotteryFeesBalanceKey, balance)
	// update the min bet value
	minBet := MustGetInt(store, minLotteryBetKey, sdk.NewIntFromUint64(math.MaxUint64))
	if bet.Amount.LT(minBet) {
		MustStoreInt(store, minLotteryBetKey, bet.Amount)
	}
	// update max bet value
	maxBet := MustGetInt(store, maxLotteryBetKey, sdk.ZeroInt())
	if bet.Amount.GT(maxBet) {
		MustStoreInt(store, maxLotteryBetKey, bet.Amount)
	}
	// we use blake 2 b to keep hashing the block hashes to compute
	// the winner of the lottery
	lh := store.Get(lotteryHashKey)
	if lh == nil {
		// init with all 0 if there is no hash yet
		lh = make([]byte, 256)
	}
	nh := blake2b.Sum256(append(lh, ctx.BlockHeader().DataHash...))
	store.Set(lotteryHashKey, nh[:])

	return nil
}

// MustGetInt gets a Int from the store panics if the int cannot be deserialized
// TODO: implement a GetInt that handle the error gracefully
func MustGetInt(store types.KVStore, key []byte, defaultValue sdk.Int) sdk.Int {
	value := store.Get(key)
	if value == nil {
		return defaultValue
	}
	var i sdk.Int
	if uErr := i.Unmarshal(store.Get(key)); uErr != nil {
		panic(fmt.Sprintf("system error - potential corrupted storage: %s", uErr.Error()))
	}
	return i
}

// MustStoreInt serialize and store an int to the store.
func MustStoreInt(store types.KVStore, key []byte, value sdk.Int) {
	binValue, err := value.Marshal()
	if err != nil {
		panic(fmt.Sprintf("system error - cannot serialize lottery data - %s", err.Error()))
	}
	store.Set(key, binValue)
}
