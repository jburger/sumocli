package remove

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdRoleRemoveUser() *cobra.Command {
	var (
		roleId string
		userId string
	)

	cmd := &cobra.Command{
		Use:   "remove user",
		Short: "Removes the specified Sumo Logic user from the role.",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role remove request started.")
			removeUserRole(roleId, userId, logger)
			logger.Debug().Msg("Role remove request finished.")
		},
	}

	cmd.Flags().StringVar(&roleId, "roleid", "", "Specify the id of the role")
	cmd.Flags().StringVar(&userId, "userid", "", "Specify the id of the user to remove")

	cmd.MarkFlagRequired("roleid")
	cmd.MarkFlagRequired("userid")
	return cmd
}

func removeUserRole(roleId string, userId string, logger zerolog.Logger) {
	client, ctx := factory.NewHttpClient()
	response, err := client.RoleManagementApi.RemoveRoleFromUser(ctx, roleId, userId)
	logging.LogError(err, logger)

	if response.StatusCode == 204 {
		fmt.Println("User: " + userId + " was removed from role: " + roleId)
	}
}
