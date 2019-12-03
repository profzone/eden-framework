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
	"github.com/profzone/eden-framework/internal/generator"
	"github.com/profzone/eden-framework/internal/generator/scanner"
	"os"

	"github.com/spf13/cobra"
)

var apiCmdInputPath, apiCmdOutputPath string

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long:  fmt.Sprintf("%s\ngenerate api doc", CommandHelpHeader),
	Run: func(cmd *cobra.Command, args []string) {
		if apiCmdInputPath == "" {
			apiCmdInputPath, _ = os.Getwd()
		}
		if apiCmdOutputPath == "" {
			apiCmdOutputPath, _ = os.Getwd()
		}
		enumScanner := scanner.NewEnumScanner()
		modelScanner := scanner.NewModelScanner(enumScanner)
		operatorScanner := scanner.NewOperatorScanner(modelScanner)
		gen := generator.NewApiGenerator(operatorScanner, modelScanner, enumScanner)

		modelScanner.Api = &gen.Api
		operatorScanner.Api = &gen.Api

		generator.Generate(gen, apiCmdInputPath, apiCmdOutputPath)
	},
}

func init() {
	generateCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&apiCmdInputPath, "input-path", "i", "", "eden generate api --input-path=/go/src/eden-server")
	apiCmd.Flags().StringVarP(&apiCmdOutputPath, "output-path", "o", "", "eden generate api --output-path=/go/src/eden-server")
}
