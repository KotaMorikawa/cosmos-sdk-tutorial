package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"hello/x/hello/types"
)

var _ = strconv.Itoa(0)

func CmdGreeting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greeting [from-address] [to-address]",
		Short: "Broadcast message greeting",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFromAddress := args[0]
			argToAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGreeting(
				clientCtx.GetFromAddress().String(),
				argFromAddress,
				argToAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
