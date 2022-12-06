package cli

import (
	"strconv"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateWhiteboard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-whiteboard [name] [description] [x-size] [y-size] [is-locked]",
		Short: "Broadcast message create-whiteboard",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argXSize, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argYSize, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argIsLocked, err := cast.ToBoolE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateWhiteboard(
				clientCtx.GetFromAddress().String(),
				argName,
				argDescription,
				argXSize,
				argYSize,
				argIsLocked,
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
