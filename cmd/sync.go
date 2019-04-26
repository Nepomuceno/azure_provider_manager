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
	"os"
	"sync"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"

	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/nepomuceno/azure_provider_manager/models"
	"github.com/spf13/cobra"
)

//var availableResourceProviders []resources.Provider
var subscriptionSync string
var inputFileName string
var outputFileName string

// initCmd represents the add command
var initCmd = &cobra.Command{
	Use:   "sync",
	Short: "Update the susbcription accoring to the profile",
	Long: `This is the update of the subscription to comply with the profile
Some sample usages would be:

azreg sync --subscription <SUBSCRIPTION_ID> --input <INPUT_FILE> 

If you want to setup the output file just pass the file path to 'output' parameter

azreg sync --subcription <SUBSCRIPTION_ID> --input <INPUT_FILE>  --output <OUPUT_FILE_PATH>
`,

	Run: func(cmd *cobra.Command, args []string) {

		providersClient = resources.NewProvidersClient(subscriptionSync)
		// create an authorizer from env vars or Azure Managed Service Idenity
		authorizer, err := auth.NewAuthorizerFromCLI()
		if err == nil {
			providersClient.Authorizer = authorizer
		} else {
			println(err)
		}
		profile := readProfileSync()
		fmt.Print(profile)
		// call the VirtualNetworks CreateOrUpdate API
	},
}

func readProfileSync() *models.Profile {

	var profile models.Profile

	if _, err := os.Stat(inputFileName); err == nil {
		jsonFile, _ := os.Open(inputFileName)
		defer jsonFile.Close()
		file, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(file, &profile)
	} else if os.IsNotExist(err) {
		fmt.Printf("Profile file not found\n")
	} else {
		fmt.Printf("%+v\n", err)
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}
	return &profile
}

func changeRegistrationForSubscription(ctx context.Context, register bool, client resources.ProvidersClient, providersToChangeRegister map[string]struct{}) error {
	var err error
	var wg sync.WaitGroup
	wg.Add(len(providersToChangeRegister))

	for providerName := range providersToChangeRegister {
		go func(p string) {
			defer wg.Done()
			fmt.Printf("[DEBUG] Registering Resource Provider %q with namespace", p)
			if register {
				if innerErr := registerWithSubscription(ctx, p, client); innerErr != nil {
					err = innerErr
				}
			} else {
				if innerErr := unregisterWithSubscription(ctx, p, client); innerErr != nil {
					err = innerErr
				}
			}
		}(providerName)
	}

	wg.Wait()

	return err
}

func unregisterWithSubscription(ctx context.Context, providerName string, client resources.ProvidersClient) error {
	if _, err := client.Unregister(ctx, providerName); err != nil {
		return fmt.Errorf("cannot un-register provider %s with Azure Resource Manager: %s", providerName, err)
	}

	return nil
}

func registerWithSubscription(ctx context.Context, providerName string, client resources.ProvidersClient) error {
	if _, err := client.Register(ctx, providerName); err != nil {
		return fmt.Errorf("cannot register provider %s with Azure Resource Manager: %s", providerName, err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&subscriptionSync, "subscription", "s", "", "Subscription id of the")
	initCmd.MarkFlagRequired("subscription")

	initCmd.Flags().StringVarP(&inputFileName, "input", "i", "", "input configuration file to be used in the sync")
	initCmd.MarkFlagRequired("input")

	initCmd.Flags().StringVarP(&outputFileName, "output", "o", "", "out configuration file tolet the generated file to get to")

}
