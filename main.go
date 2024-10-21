package main

import (
	"fmt"
	"os"
	"os/exec"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "ns" {
		fmt.Println("Usage: kubectl ns <namespace>")
		os.Exit(1)
	}

	namespace := os.Args[2]
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
	}

	// Load kubeconfig
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		fmt.Println("Error loading kubeconfig:", err)
		os.Exit(1)
	}

	// Get current context
	currentContext := config.CurrentContext
	if currentContext == "" {
		fmt.Println("No current context found in kubeconfig")
		os.Exit(1)
	}

	// Set the new namespace in the kubeconfig
	cmd := exec.Command("kubectl", "config", "set-context", currentContext, "--namespace="+namespace)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error setting namespace:", err)
		fmt.Println(string(output))
		os.Exit(1)
	}

	fmt.Printf("Namespace changed to '%s' in context '%s'\n", namespace, currentContext)
}

