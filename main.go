/*
Copyright © 2021 marcusljx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"fmt"
	"os"

	"github.com/marcusljx/ocbcctl/lib/vars"

	"github.com/marcusljx/ocbcctl/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cobra.OnInitialize(initConfig)
	cmd.Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AddConfigPath(vars.ConfigDir)
	viper.SetConfigType("yml")
	viper.SetConfigName("config")

	viper.SetDefault("callback_host", os.ExpandEnv("OCBCCTL_CALLBACK_HOST"))
	viper.SetDefault("firebase_project_id", os.ExpandEnv("FIREBASE_PROJECT_ID"))
	viper.SetDefault("firebase_collection_id", os.ExpandEnv("FIREBASE_COLLECTION_ID"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if writeErr := viper.SafeWriteConfig(); writeErr != nil {
				fmt.Printf("error writing config: %v", writeErr)
				os.Exit(2)
			}
		} else {
			fmt.Printf("unable to read config: %v\n", err)
			os.Exit(3)
		}
	}
}
