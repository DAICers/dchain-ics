package cli

import (
	"context"
	"strconv"

	"github.com/DAICers/dchain-ics/x/whiteboard/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListWhiteboardPixel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-whiteboard-pixel",
		Short: "list all whiteboard-pixel",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWhiteboardPixelRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.WhiteboardPixelAll(context.Background(), params)
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

func CmdShowWhiteboardPixel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-whiteboard-pixel [id]",
		Short: "shows a whiteboard-pixel",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetWhiteboardPixelRequest{
				Id: id,
			}

			res, err := queryClient.WhiteboardPixel(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
