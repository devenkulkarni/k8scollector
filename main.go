package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	var namespace string
	home := homedir.HomeDir()
	if home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "Kubeconfig Path")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "Kubeconfig Path")
	}

	flag.StringVar(&namespace, "namespace", "default", "Namespace to collect information from")

	flag.Parse()

	log.Printf("Using Kubeconfig %v", *kubeconfig)
	log.Printf("Collecting information for namespace %v", namespace)

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error loading kubeconfig: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	/*
		fmt.Printf("Verifying connection to cluster by fetching namespace %s...\n", namespace)
		ns, err := clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("Error fetching namespace %s: %v\n", namespace, err)
			os.Exit(1)
		}



		//fmt.Printf("Ns details: %v", ns) Gives raw data output

		// Convert it into YAML format:
		ns.Kind = "Namespace"
		ns.APIVersion = "v1"
		yamlData, err := yaml.Marshal(ns)
		if err != nil {
			fmt.Printf("Error converting to YAML: %v\n", err)
			os.Exit(1)
		}
	*/

	//fmt.Printf("NS Yaml: \n%v", string(yamlData))

	outputDir := "k8scollector-dump"
	err = os.MkdirAll(outputDir, 0755) // 0755 sets standard read/write/execute permissions
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	err = collectnamespaceyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Collection Error: %v\n", err)
	}

	err = collectpodsyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running pod yaml collector: %v", err)
	}

	err = collectdeployyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running deployment yaml collector: %v", err)
	}

	err = collectdaemonsetyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running daemonsets yaml collector: %v", err)
	}

	err = collectstsyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running Statefulsets yaml collector: %v", err)
	}

	err = collectpvcyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running PVC yaml collector: %v", err)
	}

	err = collectpvyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running PV yaml collector: %v", err)
	}

	err = collectcronjobsyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running Statefulsets yaml collector: %v", err)
	}

	err = collectingressyaml(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running Statefulsets yaml collector: %v", err)
	}

	err = collectEventsTable(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running events collector: %v", err)
	}

	err = collectPodLogs(clientset, namespace, outputDir)
	if err != nil {
		fmt.Printf("Error running pod logs collector: %v", err)
	}

}
