package cli

import (
	"strconv"

	"github.com/coreators/tokenity/x/blog/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreatePost() *cobra.Command {
	// var contents []string
	cmd := &cobra.Command{
		Use:   "create-post [contents] [description] [isStory]",
		Short: "Create a new post",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContents := args[0]
			argDescription := args[1]
			argIsStory, err := strconv.ParseBool(args[2])
			if err != nil {
				cmd.PrintErr("isStory - Boolean parsing error: ", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()

			msg := types.NewMsgCreatePost(creator, argContents, argDescription, argIsStory)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	// cmd.Flags().StringArrayVarP(&contents, "arr", "a", []string{}, "")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-post [id] [contents] [description] [isStory]",
		Short: "Update a post",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argContents := args[1]

			argDescription := args[2]
			argIsStory, err := strconv.ParseBool(args[3])
			if err != nil {
				cmd.PrintErr("isStory - Boolean parsing error: ", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePost(clientCtx.GetFromAddress().String(), id, argContents, argDescription, argIsStory)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-post [id]",
		Short: "Delete a post by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePost(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
