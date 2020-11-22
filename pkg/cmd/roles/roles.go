package roles

import (
	"github.com/spf13/cobra"
	cmdRoleList "github.com/wizedkyle/sumocli/pkg/cmd/roles/list"
)

func NewCmdRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles <command>",
		Short: "Manage roles",
		Long:  "Work with Sumo Logic roles",
	}

	//cmd.AddCommand(cmdRoleAssign.NewCmdRoleAssign())
	//cmd.AddCommand(cmdRoleCreate.NewCmdRoleCreate())
	//cmd.AddCommand(cmdRoleDelete.NewCmdRoleDelete())
	//cmd.AddCommand(cmdRoleGet.NewCmdRoleGet())
	cmd.AddCommand(cmdRoleList.NewCmdRoleList())
	//cmd.AddCommand(cmdRoleRemove.NewCmdRoleRemoveUser())
	//cmd.AddCommand(cmdRoleUpdate.NewCmdRoleUpdate())
	return cmd
}
