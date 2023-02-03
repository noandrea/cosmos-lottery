package keeper

import (
	"lottery/x/lottery/keeper"
	"testing"

	// "lottery/x/lottery"
	// "lottery/x/lottery/keeper"
	// "github.com/cosmos/cosmos-sdk/codec"
	// codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	// "github.com/cosmos/cosmos-sdk/store"
	// storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/proto/tendermint/types"
	// typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	// "github.com/stretchr/testify/require"
	// "github.com/tendermint/tendermint/libs/log"
	// tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	// tmdb "github.com/tendermint/tm-db"
	// banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func LotteryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	// storeKey := sdk.NewKVStoreKey(lottery.StoreKey)
	// memStoreKey := storetypes.NewMemoryStoreKey(lottery.MemStoreKey)

	// bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)

	// paramsStoreKey := sdk.NewKVStoreKey(paramtypes.StoreKey)
	// paramsStoreMemKey := sdk.NewKVStoreKey(paramtypes.TStoreKey)

	// db := tmdb.NewMemDB()
	// stateStore := store.NewCommitMultiStore(db)
	// stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	// stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	// require.NoError(t, stateStore.LoadLatestVersion())

	// registry := codectypes.NewInterfaceRegistry()
	// cdc := codec.NewProtoCodec(registry)

	// paramsSubspace := typesparams.NewSubspace(cdc,
	// 	lottery.Amino,
	// 	storeKey,
	// 	memStoreKey,
	// 	"LotteryParams",
	// )
	// k := keeper.NewKeeper(
	// 	cdc,
	// 	storeKey,
	// 	memStoreKey,
	// 	paramsSubspace,

	// )

	// ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// // Initialize params
	// k.SetParams(ctx, lottery.DefaultParams())

	// return k, ctx
	return nil, sdk.NewContext(nil, types.Header{}, false, nil)
}
