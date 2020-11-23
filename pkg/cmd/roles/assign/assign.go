package assign

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"strings"
)

func NewCmdRoleAssign() *cobra.Command {
	var (
		roleId string
		userId string
		output string
	)

	cmd := &cobra.Command{
		Use:   "assign",
		Short: "Assigns the specified Sumo Logic user to the role.",
		Long: `The following fields can be exported using the --output command:
name
description
filterPredicate
users
capabilities
id
`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role assign request started.")
			assignRole(roleId, userId, output, logger)
			logger.Debug().Msg("Role assign request finished.")
		},
	}

	cmd.Flags().StringVar(&roleId, "roleid", "", "Specify the id of the role")
	cmd.Flags().StringVar(&userId, "userid", "", "Specify the id of the user to remove")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")

	cmd.MarkFlagRequired("roleid")
	cmd.MarkFlagRequired("userid")

	return cmd
}

func assignRole(roleId string, userId string, output string, logger zerolog.Logger) {
	client, ctx := factory.NewHttpClient()
	roleInfo, _, err := client.RoleManagementApi.AssignRoleToUser(ctx, roleId, userId)

	logging.LogError(err, logger)

	roleInfoJson, err := json.MarshalIndent(roleInfo, "", "    ")
	logging.LogError(err, logger)

	if factory.ValidateRoleOutput(output) == true {
		value := gjson.Get(string(roleInfoJson), output)
		formattedValue := strings.Trim(value.String(), `"[]"`)
		fmt.Println(formattedValue)
	} else {
		fmt.Println(string(roleInfoJson))
	}
}
