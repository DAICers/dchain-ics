package cli

import (
	"strconv"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetWhiteboardPixelStates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-whiteboard-pixel-states [whiteboard-id]",
		Short: "Query get-whiteboard-pixel-states",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqWhiteboardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetWhiteboardPixelStatesRequest{

				WhiteboardId: reqWhiteboardId,
			}

			res, err := queryClient.GetWhiteboardPixelStates(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
