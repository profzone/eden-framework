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
	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/types"
	client "github.com/rancher/types/client/management/v3"
	pClient "github.com/rancher/types/client/project/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

// testRancherCmd represents the rancher command
var testRancherCmd = &cobra.Command{
	Use:   "testRancher",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		clientbase.Debug = false
		client, err := client.NewClient(&clientbase.ClientOpts{
			URL:       "https://localhost:8443/v3",
			AccessKey: "token-tjrrf",
			SecretKey: "lb7grj64nhhrtkg72n6hpnxr7dbqwcxtzb64b5t7d8hjqlkg5jvcfq",
			Timeout:   10 * time.Second,
			Insecure:  true,
		})
		if err != nil {
			logrus.Panicf("NewClient err: %v", err)
		}

		projects, err := client.Project.List(&types.ListOpts{})
		if err != nil {
			logrus.Panicf("Project.List err: %v", err)
		}

		p := GetProjectByName("eden", projects)

		projectClient, err := pClient.NewClient(&clientbase.ClientOpts{
			URL:       "https://localhost:8443/v3/projects/" + p.ID,
			AccessKey: "token-tjrrf",
			SecretKey: "lb7grj64nhhrtkg72n6hpnxr7dbqwcxtzb64b5t7d8hjqlkg5jvcfq",
			Timeout:   10 * time.Second,
			Insecure:  true,
		})
		if err != nil {
			logrus.Panicf("NewClient err: %v", err)
		}

		deployments, err := projectClient.Deployment.List(nil)
		if err != nil {
			logrus.Panicf("Deployment.List err: %v", err)
		}

		deployment := GetDeploymentByName("service-id", deployments)

		deployment, err = projectClient.Deployment.Update(deployment, map[string]interface{}{
			"containers": []map[string]interface{}{
				{
					"name":  "service-id",
					"image": "registry.profzone.net:5000/eden/service-id:1.0.0",
				},
			},
		})
		if err != nil {
			logrus.Panicf("Deployment.Update err: %v", err)
		}
	},
}

func GetProjectByName(name string, collection *client.ProjectCollection) *client.Project {
	for _, p := range collection.Data {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

func GetDeploymentByName(name string, collection *pClient.DeploymentCollection) *pClient.Deployment {
	for _, d := range collection.Data {
		if d.Name == name {
			return &d
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(testRancherCmd)
}