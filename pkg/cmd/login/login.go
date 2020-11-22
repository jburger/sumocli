package login

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wizedkyle/sumocli/pkg/logging"
	model "github.com/wizedkyle/sumocli/api"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Logger zerolog.Logger
func NewCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Sets Sumo Logic credentials",
		Long:  "Interactively sets the Sumo Logic Access Id, Access Key and API endpoint.",
		Run: func(cmd *cobra.Command, args []string) {
			Logger = logging.GetLoggerForCommand(cmd)
			Logger.Debug().Msg("Credential setup started.")
			configFile := configPath()
			fmt.Println("Sumocli requires an access id and access key.")
			fmt.Println("Sumocli will store the access id and access key in plain text in" +
				" the following file for use by subsequent commands:")
			fmt.Printf(configFile)
			confirmation := userConfirmation()
			if confirmation == true {
				getCredentials()
			} else {
				os.Exit(1)
			}
			Logger.Debug().Msg("Credential setup finished.")
			return
		},
	}

	return cmd
}

func configPath() string {
	var filePath string = ".sumocli/credentials/creds.json"
	homeDirectory, _ := os.UserHomeDir()
	configFile := filepath.Join(homeDirectory, filePath)
	return configFile
}

type SumoCredential struct {
	AccessId	string
	AccessKey	string
	Region		string
	Endpoint	string
}

func getCredentials() {
	sumoApiEndpoints := []model.SumoApiEndpoint{
		{Name: "Australia", Code: "au", Endpoint: "https://api.au.sumologic.com/api"},
		{Name: "Canada", Code: "ca", Endpoint: "https://api.ca.sumologic.com/api"},
		{Name: "Germany", Code: "de", Endpoint: "https://api.de.sumologic.com/api"},
		{Name: "Ireland", Code: "eu", Endpoint: "https://api.eu.sumologic.com/api"},
		{Name: "India", Code: "in", Endpoint: "https://api.in.sumologic.com/api"},
		{Name: "Japan", Code: "jp", Endpoint: "https://api.jp.sumologic.com/api"},
		{Name: "USA1", Code: "us1", Endpoint: "https://api.sumologic.com/api"},
		{Name: "USA2", Code: "us2", Endpoint: "https://api.us2.sumologic.com/api"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Name }}",
		Inactive: "{{ .Name }}",
		Selected: "{{ .Name }}",
		Details: `
------- Sumo Logic API Endpoints -------
{{ "Name:" }} {{ .Name }}
{{ "Code:" }} {{ .Code }}
{{ "Endpoint" }} {{ .Endpoint }}`,
	}

	searcher := func(input string, index int) bool {
		endpoint := sumoApiEndpoints[index]
		name := strings.Replace(strings.ToLower(endpoint.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	promptRegion := promptui.Select{
		Label:     "Please select your Sumo Logic tenancy endpoint",
		Items:     sumoApiEndpoints,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
	}

	promptAccessId := promptui.Prompt{
		Label: "Please enter your Sumo Logic Access Id",
		Mask:  '*',
	}

	promptAccessKey := promptui.Prompt{
		Label: "Please enter your Sumo Logic Access Key",
		Mask:  '*',
	}

	accessIdResult, err := promptAccessId.Run()
	accessKeyResult, err := promptAccessKey.Run()
	regionResultIndex, _, err := promptRegion.Run()
	credentials := SumoCredential{
		AccessId:  accessIdResult,
		AccessKey: accessKeyResult,
		Region:    sumoApiEndpoints[regionResultIndex].Code,
		Endpoint:  sumoApiEndpoints[regionResultIndex].Endpoint,
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	configFilePath := filepath.Dir(configPath())
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(configFilePath, 0755)
		if err != nil {
			Logger.Fatal().Err(err)
		}
	}
	credentialFile, _ := json.MarshalIndent(credentials, "", "  ")
	err = ioutil.WriteFile(configPath(), credentialFile, 0644)
	if err != nil {
		Logger.Fatal().Err(err)
	} else {
		fmt.Println("Credentials file saved to: " + configPath())
	}

	return
}

func ReadCredentialsAndEndpoint() (model.BasicAuth, string) {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(configPath()))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	accessId, accessKey, endpoint := "", "", ""
	if err != nil {
		accessId = viper.GetString("SUMO_ACCESS_ID")
		accessKey = viper.GetString("SUMO_ACCESS_KEY")
		endpoint = viper.GetString("SUMO_ENDPOINT")
	} else {
		accessId = viper.GetString("accessId")
		accessKey = viper.GetString("accessKey")
		endpoint = viper.GetString("endpoint")
	}

	return model.BasicAuth{
		UserName: accessId,
		Password: accessKey,
	}, endpoint
}

func userConfirmation() bool {
	prompt := promptui.Prompt{
		Label: "Do you want to proceed?",
	}

	result, err := prompt.Run()
	resultLower := strings.ToLower(result)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	if resultLower == "yes" {
		return true
	} else {
		fmt.Println("Error: Login cancelled")
		return false
	}
}
