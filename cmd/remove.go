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
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes chaos from the Kubernetes cluster",
	Long:  `This command removes a chaos tool installation from the Kubernetes cluster.`,
	Run: func(cmd *cobra.Command, args []string) {

		if pumba {
			var kubeConfig = getKubeConfig()
			checkKubectl(kubeConfig)
			pumbaRemove(kubeConfig)
			os.Exit(0)

		}

		if kubemonkey {
			var kubeConfig = getKubeConfig()
			checkKubectl(kubeConfig)
			kubeMonkeyRemove(kubeConfig)
			os.Exit(0)
		}

		if podreaper {
			var kubeConfig = getKubeConfig()
			checkKubectl(kubeConfig)
			podReaperRemove(kubeConfig)
			os.Exit(0)
		}

		if chaoskube {
			var kubeConfig = getKubeConfig()
			checkKubectl(kubeConfig)
			chaosKubeRemove(kubeConfig)
			os.Exit(0)
		}

		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolVarP(&pumba, "pumba", "p", false, "Remove Pumba CE tool from the cluster")
	removeCmd.Flags().BoolVarP(&podreaper, "podreaper", "r", false, "Remove pod-reaper from the cluster")
	removeCmd.Flags().BoolVarP(&chaoskube, "chaoskube", "c", false, "Remove chaoskube from the cluster")
	removeCmd.Flags().BoolVarP(&kubemonkey, "kubemonkey", "k", false, "Remove kube-monkey from the cluster")
}

func pumbaRemove(kubeConfig string) {
	if _, err := os.Stat("./manifests/pumba"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Removing chaos from Kubernetes cluster created by Pumba")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "delete", "-f", "./manifests/pumba/pumba.yml")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error removing Pumba. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully removed chaos experiments initiated by Pumba")
}

func kubeMonkeyRemove(kubeConfig string) {
	if _, err := os.Stat("./manifests/kube-monkey"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Removing chaos to Kubernetes cluster created by kube-monkey")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "delete", "-f", "./manifests/kube-monkey")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error removing kubemonkey. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully removed chaos experiment initiated by kube-monkey")
}

func podReaperRemove(kubeConfig string) {
	if _, err := os.Stat("./manifests/pod-reaper"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Removing chaos from Kubernetes cluster created by pod-reaper")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "delete", "-f", "./manifests/pod-reaper")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error removing pod-reaper. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully removed chaos experiment initiated by pod-reaper")

}

func chaosKubeRemove(kubeConfig string) {
	if _, err := os.Stat("./manifests/chaoskube"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Removing chaos to Kubernetes cluster created by chaoskube")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "delete", "-f", "./manifests/chaoskube")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error removing chaoskube. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully removed chaos experiment initiated by chaoskube")
}
