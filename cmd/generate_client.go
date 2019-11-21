/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"profzone/eden-framework/internal/generator"

	"github.com/spf13/cobra"
)

var (
	clientCmdApiPath, clientCmdPackageName string
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long:  fmt.Sprintf("%s\ngenerate client", CommandHelpHeader),
	Run: func(cmd *cobra.Command, args []string) {
		gen := &generator.ClientGenerator{
			ServiceName: clientCmdPackageName,
		}
		generator.Generate(gen, clientCmdApiPath)
	},
}

func init() {
	generateCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVarP(&clientCmdPackageName, "package-name", "p", "", "eden generate api --package-name=client_account")
	clientCmd.Flags().StringVarP(&clientCmdApiPath, "api-path", "f", "", "eden generate api --api-path=/go/src/eden-server/api.json")
}
