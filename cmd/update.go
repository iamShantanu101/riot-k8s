// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"bytes"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the chaos engineering experiments' attributes",
	Long:  `This command is used to update the attributes of the ongoing chaos engineering experiment.`,
	Run: func(cmd *cobra.Command, args []string) {

		if pumba {
			var kubeConfig = getKubeConfig()
			checkKubectl(kubeConfig)
			pumbaUpdate(kubeConfig)
			os.Exit(0)

		}

		if kubemonkey {
			var kubeConfig = getKubeConfig()
			checkKubectl(kubeConfig)
			kubeMonkeyUpdate(kubeConfig)
			os.Exit(0)
		}

		if podreaper {
			var kubeConfig = getKubeConfig()
			checkKubectl(kubeConfig)
			podReaperUpdate(kubeConfig)
			os.Exit(0)
		}

		if chaoskube {
			var kubeConfig = getKubeConfig()
			checkKubectl(kubeConfig)
			chaosKubeUpdate(kubeConfig)
			os.Exit(0)
		}

		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVarP(&pumba, "pumba", "p", false, "Updates the Pumba CE config on the cluster")
	updateCmd.Flags().BoolVarP(&kubemonkey, "kubemonkey", "k", false, "Updates kube-monkey config on the cluster")
	updateCmd.Flags().BoolVarP(&podreaper, "podreaper", "r", false, "Updates pod-reaper config on the cluster")
	updateCmd.Flags().BoolVarP(&chaoskube, "chaoskube", "c", false, "Updates chaoskube configon the cluster")

}

func pumbaUpdate(kubeConfig string) {
	if _, err := os.Stat("./manifests/pumba"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Updating chaos in Kubernetes cluster created by Pumba")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "apply", "-f", "./manifests/pumba/pumba.yml")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error updating Pumba. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully updated chaos experiments initiated by Pumba")
}

func kubeMonkeyUpdate(kubeConfig string) {
	if _, err := os.Stat("./manifests/kube-monkey"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Updating chaos in Kubernetes cluster created by kube-monkey")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "apply", "-f", "./manifests/kube-monkey")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error updating kube-monkey. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully updated chaos experiment initiated by kube-monkey")
}

func podReaperUpdate(kubeConfig string) {
	if _, err := os.Stat("./manifests/pod-reaper"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Updating chaos in Kubernetes cluster created by pod-reaper")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "apply", "-f", "./manifests/pod-reaper")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error updating pod-reaper. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully updated chaos experiment initiated by pod-reaper")

}

func chaosKubeUpdate(kubeConfig string) {
	if _, err := os.Stat("./manifests/chaoskube"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Updating chaos in Kubernetes cluster created by chaoskube")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "apply", "-f", "./manifests/chaoskube")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error updating chaoskube. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully updated chaos experiment initiated by chaoskube")
}
