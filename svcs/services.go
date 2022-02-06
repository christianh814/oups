package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// used to interact with services in .local/share/oups

var ValidPluginFilenamePrefixes = []string{"oups"}
var configs = make(map[string]viper.Viper)

func Run() {
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
