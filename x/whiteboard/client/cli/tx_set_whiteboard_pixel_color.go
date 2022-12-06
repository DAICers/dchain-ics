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

func CmdSetWhiteboardPixelColor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-whiteboard-pixel-color [whiteboard-id] [x] [y] [r] [g] [b]",
		Short: "Broadcast message set-whiteboard-pixel-color",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argWhiteboardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argX, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argY, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argR, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argG, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argB, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetWhiteboardPixelColor(
				clientCtx.GetFromAddress().String(),
				argWhiteboardId,
				argX,
				argY,
				argR,
				argG,
				argB,
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
