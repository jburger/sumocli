package list

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	sl "github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"github.com/wizedkyle/sumocli/pkg/options"
	"strings"
)

func NewCmdRoleList() *cobra.Command {
	var (
		numberOfResults string
		filter          string
		output          string
		token			string
		sort			string
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists Sumo Logic roles",
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
			logger.Debug().Msg("Role list request started.")
			listRoles(numberOfResults, filter, token, sort, output, logger)
			logger.Debug().Msg("Role list request finished.")
		},
	}

	cmd.Flags().StringVar(&numberOfResults, "results", "", "Specify the number of results, this is set to 100 by default.")
	cmd.Flags().StringVar(&filter, "filter", "", "Specify the name of the role you want to retrieve")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from")
	cmd.Flags().StringVar(&token, "token", "", "Specify the token")
	cmd.Flags().StringVar(&sort, "sort", "", "Specify the field to sort by")

	return cmd
}


func listRoles(numberOfResults string, filter string, token string, sort string, output string, logger zerolog.Logger) {
	client, ctx := factory.NewHttpClient()
	roles, response, err2 := client.RoleManagementApi.ListRoles(ctx, &sl.RoleManagementApiListRolesOpts{
		Limit:  options.StringToInt32(numberOfResults),
		Token:  options.StringToOptional(token),
		SortBy: options.StringToOptional(sort),
		Name: options.StringToOptional(filter),
	})
	logging.LogError(err2, logger)
	if response.StatusCode != 200 || err2 != nil {
		return
	}

	json, jsonErr := json.MarshalIndent(roles.Data, "", "    ")
	logging.LogError(jsonErr, logger)

	if factory.ValidateRoleOutput(output) == true {
		value := gjson.Get(string(json), "#."+output)
		filteredOutput := strings.Trim(value.String(), `"[]"`)
		fmt.Println(filteredOutput)
	} else {
		fmt.Println(string(json))
	}
}
