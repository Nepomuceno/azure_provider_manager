// Copyright © 2019 Gabriel Nepomuceno
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
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var unsupportedResources = []string{"AppDynamics.APM",
"LiveArena.Broadcast",
"microsoft.aadiam",
"Microsoft.BizTalkServices",
"Microsoft.ContentModerator",
"Microsoft.CustomerLockbox",
"Microsoft.DynamicsLcs",
"Microsoft.EnterpriseKnowledgeGraph",
"Microsoft.HanaOnAzure",
"Microsoft.Intune",
"Microsoft.MachineLearningModelManagement",
"Microsoft.StorSimple",
"microsoft.visualstudio",
"Microsoft.WindowsDefenderATP",
"Pokitdok.Platform",
"TrendMicro.DeepSecurity",
}

func supportedResource(e string) bool {
    for _, a := range unsupportedResources {
        if a == e {
            return false
        }
    }
    return true
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azreg",
	Short: "Manage Azure registry",
	Long: `Azure Registry manager it is a command line tool
	that allows you to registry azure provider or unregistry them`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	
	
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Gabriel Nepomuceno <ganepomuceno@microsoft.com>")
	viper.SetDefault("license", "MIT")
	viper.SetDefault("version", "v1.0.0")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
