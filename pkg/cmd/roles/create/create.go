package create

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io/ioutil"
	"os"
)

func NewCmdRoleCreate() *cobra.Command {
	var (
		name         string
		description  string
		filter       string
		users        []string
		capabilities []string
		autofill     bool
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a Sumo Logic role",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLoggerForCommand(cmd)
			logger.Debug().Msg("Role create request started.")
			createRole(name, description, filter, users, capabilities, autofill, logger)
			logger.Debug().Msg("Role create request finished.")
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Name of the role")
	cmd.Flags().StringVar(&description, "description", "", "Description for the role")
	cmd.Flags().StringVar(&filter, "filter", "", "Search filter for the role")
	cmd.Flags().StringSliceVar(&users, "users", []string{}, "Comma deliminated list of user ids to add to the role")
	cmd.Flags().StringSliceVar(&capabilities, "capabilities", []string{}, "Comma deliminated list of capabilities")
	cmd.Flags().BoolVar(&autofill, "autofill", true, "Is set to true by default.")

	return cmd
}

func createRole(name string, description string, filter string, users []string, capabilities []string, autofill bool, logger zerolog.Logger) {
	var createRoleResponse api.RoleData

	if name == "" {
		fmt.Println("--name field needs to be specified.")
		os.Exit(0)
	}

	for i, capability := range capabilities {
		if validateCapabilities(capability) == false {
			fmt.Println(capability + " is not a valid Sumo Logic role capability.")
		}
		i++
	}

	requestBodySchema := &api.CreateRoleRequest{
		Name:                 name,
		Description:          description,
		FilterPredicate:      filter,
		Users:                users,
		Capabilities:         capabilities,
		AutoFillDependencies: autofill,
	}

	requestBody, _ := json.Marshal(requestBodySchema)
	client, request := factory.NewHttpRequestWithBody("POST", "v1/roles", requestBody)
	response, err := client.Do(request)
	logging.LogError(err, logger)

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, logger)
	} else {
		jsonErr := json.Unmarshal(responseBody, &createRoleResponse)
		logging.LogError(jsonErr, logger)
		fmt.Println(createRoleResponse.Name + " role successfully created")
	}
}

func validateCapabilities(capability string) bool {
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
