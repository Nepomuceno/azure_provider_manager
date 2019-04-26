// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"

	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/nepomuceno/azure_provider_manager/models"
	"github.com/spf13/cobra"
)

var subscription string
var profileName string
var outputFile string
var providersClient resources.ProvidersClient

// initCmd represents the init command
var addCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `This is the initializaiton of the profile
Some sample usages would be:

azreg init --profile <PROFILE_NAME> --subscription <SUBSCRIPTION_ID>

If you want to setup the output file just pass the file path to 'output' parameter

azreg init --profile <PROFILE_NAME> --subcription <SUBSCRIPTION_ID> --output <ABSOLUTE_FILE_PATH>
`,

	Run: func(cmd *cobra.Command, args []string) {

		providersClient = resources.NewProvidersClient(subscription)
		// create an authorizer from env vars or Azure Managed Service Idenity
		authorizer, err := auth.NewAuthorizerFromCLI()
		if err == nil {
			providersClient.Authorizer = authorizer
		} else {
			println(err)
		}

		fmt.Printf("Init profile: %s\n", profileName)
		profile := readProfile()
		fmt.Print(profile)
	},
}

func readProfile() *models.Profile {

	var profile models.Profile
	if outputFile == "" {

	}
	fileName := fmt.Sprintf("./%s.profile.json", profileName)

	providersList, err := providersClient.List(context.Background(), nil, "")
	var enabled []string
	var disabled []string
	if err == nil {
		providersValues := providersList.Values()
		for _, prov := range providersValues {
			if supportedResource(*prov.Namespace) {
				if *prov.RegistrationState == "NotRegistered" {
					disabled = append(disabled, *prov.Namespace)
				}
				if *prov.RegistrationState == "Registered" {
					enabled = append(enabled, *prov.Namespace)
				}
			}
		}
	}
	file := &models.Profile{profileName, enabled, disabled}
	profile = *file
	content, _ := json.Marshal(file)
	ioutil.WriteFile(fileName, content, 0644)
	fmt.Println(string(content))
	return &profile
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addCmd.Flags().StringVarP(&subscription, "subscription", "s", "", "Subscription id of the")
	addCmd.MarkFlagRequired("subscription")
	addCmd.Flags().StringVarP(&profileName, "profile", "p", "default", "profile to be used")
	addCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to put the data on")
}
