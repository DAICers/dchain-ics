package cli

import (
	"context"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListWhiteboardPixelMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-whiteboard-pixel-map",
		Short: "list all whiteboard-pixel-map",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWhiteboardPixelMapRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.WhiteboardPixelMapAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowWhiteboardPixelMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-whiteboard-pixel-map [whiteboard-id] [index]",
		Short: "shows a whiteboard-pixel-map",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argWhiteboardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argIndex, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryGetWhiteboardPixelMapRequest{
				WhiteboardId: argWhiteboardId,
				Index:        argIndex,
			}

			res, err := queryClient.WhiteboardPixelMap(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
