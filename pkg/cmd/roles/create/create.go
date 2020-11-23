package create

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"strings"
)

func NewCmdRoleCreate() *cobra.Command {
	var (
		name         string
		description  string
		filter       string
		users        []string
		capabilities []string
		autofill     bool
		output       string
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role create request started.")
			createRole(name, description, filter, users, capabilities, autofill, output, logger)
			logger.Debug().Msg("Role create request finished.")
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Name of the role.")
	cmd.Flags().StringVar(&description, "description", "", "Description for the role.")
	cmd.Flags().StringVar(&filter, "filter", "", "Search filter for the role.")
	cmd.Flags().StringSliceVar(&users, "users", []string{}, "Comma deliminated list of user ids to add to the role.")
	cmd.Flags().StringSliceVar(&capabilities, "capabilities", []string{}, "Comma deliminated list of capabilities.")
	cmd.Flags().BoolVar(&autofill, "autofill", true, "Is set to true by default.")
	cmd.Flags().StringVar(&output, "output", "", "Specify the field to export the value from.")

	cmd.MarkFlagRequired("name")

	return cmd
}

func createRole(name string, description string, filter string, users []string, capabilities []string, autofill bool, output string, logger zerolog.Logger) {
	if !validateCapabilities(capabilities) {
		return
	}

	client, ctx := factory.NewHttpClient()
	roleInfo, _, err := client.RoleManagementApi.CreateRole(ctx, api.CreateRoleDefinition{
		Name:                 name,
		Description:          description,
		FilterPredicate:      filter,
		Users:                users,
		Capabilities:         capabilities,
		AutofillDependencies: autofill,
	})

	logging.LogError(err, logger)

	createRoleResponseJson, err := json.MarshalIndent(roleInfo, "", "    ")
	logging.LogError(err, logger)

	if factory.ValidateRoleOutput(output) == true {
		value := gjson.Get(string(createRoleResponseJson), output)
		formattedValue := strings.Trim(value.String(), `"[]"`)
		fmt.Println(formattedValue)
	} else {
		fmt.Println(string(createRoleResponseJson))
		fmt.Println(roleInfo.Name + " role successfully created")
	}
}

func validateCapabilities(capabilities []string) bool {
	success := true
	for i, capability := range capabilities {
		if validateCapability(capability) == false {
			fmt.Println(capability + " is not a valid Sumo Logic role capability.")
			success = false
		}
		i++
	}
	return success
}

func validateCapability(capability string) bool {
	switch capability {
	case
		"viewCollectors",
		"manageCollectors",
		"manageBudgets",
		"manageDataVolumeFeed",
		"viewFieldExtraction",
		"manageFieldExtractionRules",
		"manageS3DataForwarding",
		"manageContent",
		"dataVolumeIndex",
		"viewConnections",
		"manageConnections",
		"viewScheduledViews",
		"manageScheduledViews",
		"viewPartitions",
		"managePartitions",
		"viewFields",
		"manageFields",
		"viewAccountOverview",
		"manageTokens",
		"manageDataStreams",
		"manageEntityTypeConfig",
		"manageMonitors",
		"metricsTransformation",
		"metricsExtraction",
		"metricsRules",
		"managePasswordPolicy",
		"ipWhitelisting",
		"createAccessKeys",
		"manageAccessKeys",
		"manageSupportAccountAccess",
		"manageAuditDataFeed",
		"manageSaml",
		"shareDashboardOutsideOrg",
		"manageOrgSettings",
		"changeDataAccessLevel",
		"shareDashboardWorld",
		"shareDashboardWhitelist",
		"manageUsersAndRoles",
		"searchAuditIndex",
		"auditEventIndex":
		return true
	}
	return false
}
