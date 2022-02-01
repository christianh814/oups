/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  pluginCmdLongText,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		run()
	},
}

var ValidPluginFilenamePrefixes = []string{"oups"}
var configs = make(map[string]viper.Viper)

func init() {
	pluginCmd.AddCommand(listCmd)
}

func run() {
	pluginsFound := false
	isFirstFile := true
	dir := "/home/robert/.local/share/oups"

	files, err := ioutil.ReadDir(dir)
	fmt.Println(err)

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if !hasValidPrefix(f.Name(), ValidPluginFilenamePrefixes) {
			continue
		}

		if isFirstFile {
			fmt.Println("The following compatible plugins are available:\n\n")
			pluginsFound = true
			isFirstFile = false
		}

		pluginPath := f.Name()
		//		if !o.NameOnly {
		pluginPath = filepath.Join(dir, pluginPath)
		//		}

		fmt.Println(pluginPath)
	}

	if !pluginsFound {
		fmt.Println("no plugins found")
	}
	readConfig()
}
func hasValidPrefix(filepath string, validPrefixes []string) bool {
	for _, prefix := range validPrefixes {
		if !strings.HasPrefix(filepath, prefix+"-") {
			continue
		}
		return true
	}
	return false
}
func readConfig() {

	// TODO: this needs to be generic to read in all configs
	loadbalancer := viper.New()
	loadbalancer.AddConfigPath("/home/robert/.local/share/oups/")
	loadbalancer.SetConfigName("oups-loadbalancer")
	loadbalancer.SetConfigType("yaml")
	loadbalancer.ReadInConfig()
	fmt.Println(loadbalancer.Get("image"))
	configs["loadbalancer"] = *loadbalancer

}
