package delete

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
)

func NewCmdRoleDelete() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role delete request started.")
			deleteRole(id, logger)
			logger.Debug().Msg("Role delete request finished.")
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the role to delete")

	return cmd
}

func deleteRole(id string, logger zerolog.Logger) {
	client, ctx := factory.NewHttpClient()
	role, _, err := client.RoleManagementApi.GetRole(ctx, id)
	logging.LogError(err, logger)
	if len(role.Users) > 0 {
		logger.Warn().Msg("The role wasn't deleted. Users have been assigned to it. " +
			"Please run `sumocli remove user` then re-run `sumocli roles delete`.")
	}
	response, err := client.RoleManagementApi.DeleteRole(ctx, id)
	logging.LogError(err, logger)

	if response.StatusCode == 204 {
		logger.Info().Msg("Role was deleted.")
	}
}