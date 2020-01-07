/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ciRunCmdInDocker bool

// ciRunCmd represents the run command
var ciRunCmd = &cobra.Command{
	Use:   "run",
	Short: "ci run project's script",
	Long:  fmt.Sprintf("%s\nci run project's script", CommandHelpHeader),
}

func init() {
	ciCmd.AddCommand(ciRunCmd)
	ciRunCmd.Flags().BoolVarP(&ciRunCmdInDocker, "in-docker", "d", false, "run script in docker")

	if currentProject.Scripts != nil {
		for scriptCmd, script := range currentProject.Scripts {
			ciRunCmd.AddCommand(&cobra.Command{
				Use:   scriptCmd,
				Short: script.String(),
				Run: func(cmd *cobra.Command, args []string) {
					err := currentProject.RunScript(cmd.Use, ciRunCmdInDocker)
					if err != nil {
						logrus.Error(err)
					}
				},
			})
		}
	}
}
