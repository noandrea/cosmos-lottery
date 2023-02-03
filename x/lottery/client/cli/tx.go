package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"lottery/x/lottery"

	"github.com/cosmos/cosmos-sdk/version"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        lottery.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", lottery.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(NewPlaceBet())

	return cmd
}

// egTx utility to generate example command string
func egTx(cmd ...string) string {
	return fmt.Sprintln(version.AppName, "tx", lottery.ModuleName, strings.Join(cmd, " "))
}

// NewPlaceBet defines the command to publish credentials
func NewPlaceBet() *cobra.Command {

	var (
		command = "place-bet"
	)

	cmd := &cobra.Command{
		Use:     fmt.Sprintln(command, "bet-amount"),
		Short:   "place a bet for the current lottery",
		Example: egTx(command, "10"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// parse the amount
			amount, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("the bet-amount argument must be an int number")
			}
			// retrieve the signer
			signer := clientCtx.GetFromAddress()
			// create the message
			msg := lottery.NewMsgPlaceBetRequest(
				sdk.NewInt64Coin(sdk.DefaultBondDenom, amount),
				signer,
			)
			// execute
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// add flags
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
