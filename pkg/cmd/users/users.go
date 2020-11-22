package users

import (
	"github.com/spf13/cobra"
)

func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users <command>",
		Short: "Manage users",
		Long:  "Work with Sumo Logic users",
	}

	//cmd.AddCommand(cmdUserChange.NewCmdUserChangeEmail())
	//cmd.AddCommand(cmdUserCreate.NewCmdUserCreate())
	//cmd.AddCommand(cmdUserDelete.NewCmdUserDelete())
	//cmd.AddCommand(cmdUserGet.NewCmdGetUser())
	//cmd.AddCommand(cmdUserList.NewCmdUserList())
	//cmd.AddCommand(cmdUserUnlock.NewCmdUnlockUser())
	return cmd
}
