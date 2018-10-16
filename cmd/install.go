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
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var pumba, chaoskube, kubemonkey, podreaper bool

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a chaos engineering tool on Kubernetes cluster",
	Long:  `Used for installing a chaos engineering tool on K8S cluster.`,
	Run: func(cmd *cobra.Command, args []string) {

		if pumba {
			fmt.Println("inside pumba")
			kubeConfig := getKubeConfig()
			checkKubectl(kubeConfig)
			pumbaInstall(kubeConfig)
			os.Exit(0)
		}

		if kubemonkey {
			kubeConfig := getKubeConfig()
			checkKubectl(kubeConfig)
			kubeMonkeyInstall(kubeConfig)
			os.Exit(0)
		}

		if podreaper {
			kubeConfig := getKubeConfig()
			checkKubectl(kubeConfig)
			podReaperInstall(kubeConfig)
			os.Exit(0)
		}

		if chaoskube {
			kubeConfig := getKubeConfig()
			checkKubectl(kubeConfig)
			chaosKubeInstall(kubeConfig)
			os.Exit(0)
		}
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&pumba, "pumba", "p", false, "Deploy Pumba CE tool on the cluster")
	installCmd.Flags().BoolVarP(&kubemonkey, "kubemonkey", "k", false, "Deploy kube-monkey on the cluster")
	installCmd.Flags().BoolVarP(&podreaper, "podreaper", "r", false, "Deploy pod-reaper on the cluster")
	installCmd.Flags().BoolVarP(&chaoskube, "chaoskube", "c", false, "Deploy chaoskube on the cluster")
}

func getKubeConfig() string {
	// Get kubeconfig
	fmt.Println("Please enter the path to your kubeconfig:")
	var kubeConfig string
	fmt.Scanln(&kubeConfig)
	fmt.Printf("path: %s\n", kubeConfig)
	if _, err := os.Stat(kubeConfig); err != nil {
		fmt.Println("Kubeconfig file not found, kindly check")
		os.Exit(1)
	}
	return kubeConfig
}

func checkKubectl(kubeConfig string) {
	/*This function is used to check the whether kubectl command is installed &
	  it works with the kubeConfig provided
	*/
	kctldir, err := exec.LookPath("kubectl")
	if err != nil {
		log.Fatal("kubectl command not found. Please check if kubectl is installed")
	}
	fmt.Printf("Found kubectl at %s\n", kctldir)
	kver, err := exec.Command("kubectl", "--kubeconfig", kubeConfig, "version", "--short").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(kver))
}

func pumbaInstall(kubeConfig string) {
	if _, err := os.Stat("./manifests/pumba"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Adding chaos to Kubernetes cluster with Pumba")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "create", "-f", "./manifests/pumba/pumba.yml")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error initiating chaos experiment with Pumba. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully initiated chaos experiment with Pumba")
}

func kubeMonkeyInstall(kubeConfig string) {
	if _, err := os.Stat("./manifests/kube-monkey"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Adding chaos to Kubernetes cluster with kube-monkey")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "create", "-f", "./manifests/kube-monkey")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error initiating chaos experiment with kubemonkey. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully initiated chaos experiment with kube-monkey")
}

func podReaperInstall(kubeConfig string) {
	if _, err := os.Stat("./manifests/pod-reaper"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Adding chaos to Kubernetes cluster with pod-reaper")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "create", "-f", "./manifests/pod-reaper")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error initiating chaos experiment with pod-reaper. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully initiated chaos experiment with pod-reaper")

}

func chaosKubeInstall(kubeConfig string) {
	if _, err := os.Stat("./manifests/chaoskube"); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	fmt.Println("Adding chaos to Kubernetes cluster with chaoskube")
	chaos := exec.Command("kubectl", "--kubeconfig", kubeConfig, "create", "-f", "./manifests/chaoskube")
	var stdout, stderr bytes.Buffer
	chaos.Stdout = &stdout
	chaos.Stderr = &stderr
	err := chaos.Run()
	if err != nil {

		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Error initiating chaos experiment with chaoskube. Please check the configuration")
		log.Fatalf("chaos.Run() failed with %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully initiated chaos experiment with chaoskube")
}
