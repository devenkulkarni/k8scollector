package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/yaml"
)

func collectnamespaceyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {
	fmt.Printf("Fetching namespace %s yaml\n", namespace)
	ns, err := clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error fetching namespace %s: %v\n", namespace, err)
		os.Exit(1)
	}

	ns.Kind = "Namespace"
	ns.APIVersion = "v1"
	yamlData, err := yaml.Marshal(ns)
	if err != nil {
		fmt.Printf("Error converting to YAML: %v\n", err)
		os.Exit(1)
	}

	filePath := filepath.Join(outputDir, "namespace.yaml")
	err = os.WriteFile(filePath, yamlData, 0644) // 0644 sets standard file permissions
	if err != nil {
		fmt.Printf("Error writing file to disk: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully saved clean namespace config to: %s\n", filePath)
	return nil
}

func collectpodsyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {

	fmt.Printf("Fetching all pod yamls in namespace - %s \n", namespace)
	podlist, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching pods in namespace - %s: %v", namespace, err)
	}

	podsDir := filepath.Join(outputDir, "pods")
	err = os.MkdirAll(podsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create pods directory: %w", err)
	}

	for _, pods := range podlist.Items {
		pods.Kind = "Pod"
		pods.APIVersion = "v1"
		podyamldata, err := yaml.Marshal(pods)
		if err != nil {
			fmt.Printf("Error converting to YAML: %v", err)
			continue
		}

		filePath := filepath.Join(podsDir, pods.Name+".yaml")
		err = os.WriteFile(filePath, podyamldata, 0644) // 0644 sets standard file permissions
	}
	if err != nil {
		fmt.Printf("Error writing file to disk: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nSaved all pods yamls in namespace - %s in directory - %v\n", namespace, podsDir)
	return nil
}

func collectdeployyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {

	fmt.Printf("Fetching all Deployment yamls in namespace - %s \n", namespace)
	deploylist, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching deployments in namespace - %s: %v", namespace, err)
	}

	deployDir := filepath.Join(outputDir, "deployments")
	err = os.MkdirAll(deployDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create deployments directory: %w", err)
	}

	for _, deployments := range deploylist.Items {
		deployments.Kind = "Deployment"
		deployments.APIVersion = "apps/v1"
		deploymentsyamldata, err := yaml.Marshal(deployments)
		if err != nil {
			fmt.Printf("Error converting to YAML: %v", err)
			continue
		}

		filePath := filepath.Join(deployDir, deployments.Name+".yaml")
		err = os.WriteFile(filePath, deploymentsyamldata, 0644) // 0644 sets standard file permissions
	}
	if err != nil {
		fmt.Printf("Error writing file to disk: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nSaved all Deployment yamls in namespace - %s in directory - %v\n", namespace, deployDir)
	return nil
}

func collectdaemonsetyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {

	fmt.Printf("Fetching all Daemonsets yamls in namespace - %s \n", namespace)
	daemonlist, err := clientset.AppsV1().DaemonSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching daemonsets in namespace - %s: %v", namespace, err)
	}

	daemonDir := filepath.Join(outputDir, "daemonsets")
	err = os.MkdirAll(daemonDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create daemonsets directory: %w", err)
	}

	for _, daemonsets := range daemonlist.Items {
		daemonsets.Kind = "DaemonSet"
		daemonsets.APIVersion = "apps/v1"
		daemonsetsyamldata, err := yaml.Marshal(daemonsets)
		if err != nil {
			fmt.Printf("Error converting to YAML: %v", err)
			continue
		}

		filePath := filepath.Join(daemonDir, daemonsets.Name+".yaml")
		err = os.WriteFile(filePath, daemonsetsyamldata, 0644) // 0644 sets standard file permissions
	}
	if err != nil {
		fmt.Printf("Error writing file to disk: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nSaved all Daemonsets yamls in namespace - %s in directory - %v\n", namespace, daemonDir)
	return nil
}

func collectstsyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {

	fmt.Printf("Fetching all Statefulsets yamls in namespace - %s \n", namespace)
	sts, err := clientset.AppsV1().StatefulSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching Statefulsets in namespace - %s: %v", namespace, err)
	}

	stsDir := filepath.Join(outputDir, "statefulsets")
	err = os.MkdirAll(stsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create statefulsets directory: %w", err)
	}

	for _, stsitems := range sts.Items {
		stsitems.Kind = "Deployment"
		stsitems.APIVersion = "apps/v1"
		stsyamldata, err := yaml.Marshal(stsitems)
		if err != nil {
			fmt.Printf("Error converting to YAML: %v", err)
			continue
		}

		filePath := filepath.Join(stsDir, stsitems.Name+".yaml")
		err = os.WriteFile(filePath, stsyamldata, 0644) // 0644 sets standard file permissions
	}
	if err != nil {
		fmt.Printf("Error writing file to disk: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nSaved all Statefulsets yamls in namespace - %s in directory - %v\n", namespace, stsDir)
	return nil
}

func collectpvcyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {

	fmt.Printf("Fetching all PVC yamls in namespace - %s \n", namespace)
	pvclist, err := clientset.CoreV1().PersistentVolumeClaims(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching PVC in namespace - %s: %v", namespace, err)
	}

	pvcDir := filepath.Join(outputDir, "persistentvolumeclaims")
	err = os.MkdirAll(pvcDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create pvc directory: %w", err)
	}

	for _, pvc := range pvclist.Items {
		pvc.Kind = "PersistenVolumeClaim"
		pvc.APIVersion = "v1"
		pvcyamldata, err := yaml.Marshal(pvc)
		if err != nil {
			fmt.Printf("Error converting to YAML: %v", err)
			continue
		}

		filePath := filepath.Join(pvcDir, pvc.Name+".yaml")
		err = os.WriteFile(filePath, pvcyamldata, 0644) // 0644 sets standard file permissions
	}
	if err != nil {
		fmt.Printf("Error writing file to disk: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nSaved all PVC yamls in namespace - %s in directory - %v\n", namespace, pvcDir)
	return nil
}

func collectpvyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {

	fmt.Printf("Fetching all PV yamls that refer to PVCs in namespace - %s \n", namespace)
	pvclist, err := clientset.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching PV in namespace - %s: %v", namespace, err)
	}

	pvDir := filepath.Join(outputDir, "persistenvolumes")
	err = os.MkdirAll(pvDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create pvc directory: %w", err)
	}

	for _, pv := range pvclist.Items {
		if pv.Spec.ClaimRef == nil || pv.Spec.ClaimRef.Namespace != namespace {
			pv.Kind = "PersistenVolume"
			pv.APIVersion = "v1"
			pvyamldata, err := yaml.Marshal(pv)
			if err != nil {
				fmt.Printf("Error converting to YAML: %v", err)
				continue
			}

			filePath := filepath.Join(pvDir, pv.Name+".yaml")
			err = os.WriteFile(filePath, pvyamldata, 0644) // 0644 sets standard file permissions
		}
		if err != nil {
			fmt.Printf("Error writing file to disk: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("\nSaved all PV yamls in namespace - %s in directory - %v\n", namespace, pvDir)
	return nil
}

func collectcronjobsyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {

	fmt.Printf("Fetching all Cronjobs yamls in namespace - %s \n", namespace)
	cjlist, err := clientset.BatchV1().CronJobs(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching Cronjobs in namespace - %s: %v", namespace, err)
	}

	cjDir := filepath.Join(outputDir, "cronjobs")
	err = os.MkdirAll(cjDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create cronjobs directory: %w", err)
	}

	for _, cronjobs := range cjlist.Items {
		cronjobs.Kind = "CronJob"
		cronjobs.APIVersion = "batch/v1"
		cronjobsyamldata, err := yaml.Marshal(cronjobs)
		if err != nil {
			fmt.Printf("Error converting to YAML: %v", err)
			continue
		}

		filePath := filepath.Join(cjDir, cronjobs.Name+".yaml")
		err = os.WriteFile(filePath, cronjobsyamldata, 0644) // 0644 sets standard file permissions
	}
	if err != nil {
		fmt.Printf("Error writing file to disk: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nSaved all Cronjobs yamls in namespace - %s in directory - %v\n", namespace, cjDir)
	return nil
}

func collectingressyaml(clientset *kubernetes.Clientset, namespace string, outputDir string) error {

	fmt.Printf("Fetching all Cronjobs yamls in namespace - %s \n", namespace)
	inglist, err := clientset.NetworkingV1().Ingresses(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error fetching Ingresses in namespace - %s: %v", namespace, err)
	}

	ingDir := filepath.Join(outputDir, "ingress")
	err = os.MkdirAll(ingDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create ingress directory: %w", err)
	}

	for _, ingress := range inglist.Items {
		ingress.Kind = "CronJob"
		ingress.APIVersion = "batch/v1"
		ingressyamldata, err := yaml.Marshal(ingress)
		if err != nil {
			fmt.Printf("Error converting to YAML: %v", err)
			continue
		}

		filePath := filepath.Join(ingDir, ingress.Name+".yaml")
		err = os.WriteFile(filePath, ingressyamldata, 0644) // 0644 sets standard file permissions
	}
	if err != nil {
		fmt.Printf("Error writing file to disk: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\nSaved all Ingress yamls in namespace - %s in directory - %v\n", namespace, ingDir)
	return nil
}

func collectEventsTable(clientset *kubernetes.Clientset, namespace string, outputDir string) error {
	fmt.Printf("Fetching and formatting events in namespace: %s\n", namespace)

	// 1. Fetch the raw events
	eventList, err := clientset.CoreV1().Events(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to fetch events: %w", err)
	}

	// 2. Initialize our log text string with the exact header columns
	// The numbers specify column spacing width
	tableContent := fmt.Sprintf("%-10s %-10s %-20s %-50s %s\n", "LAST SEEN", "TYPE", "REASON", "OBJECT", "MESSAGE")

	now := time.Now()

	// 3. Loop over the events and construct rows
	for _, event := range eventList.Items {
		// Calculate a friendly "Last Seen" duration string
		lastSeenStr := "unknown"
		if !event.LastTimestamp.IsZero() {
			duration := now.Sub(event.LastTimestamp.Time).Round(time.Second)
			lastSeenStr = duration.String() // Converts to strings like "3s", "1m15s"
		}

		// Extract the object reference (e.g., "pod/error-deploy-xyz")
		objectRef := fmt.Sprintf("%s/%s", strings.ToLower(event.InvolvedObject.Kind), event.InvolvedObject.Name)

		// Format this single row to match the header widths perfectly
		row := fmt.Sprintf("%-10s %-10s %-20s %-50s %s\n",
			lastSeenStr,
			event.Type,
			event.Reason,
			objectRef,
			event.Message,
		)

		tableContent += row
	}

	filePath := filepath.Join(outputDir, "events.log")
	err = os.WriteFile(filePath, []byte(tableContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to save events table: %w", err)
	}

	fmt.Printf("✓ Saved formatted events log to: %s\n", filePath)
	return nil
}

func collectPodLogs(clientset *kubernetes.Clientset, namespace string, outputDir string) error {
	fmt.Printf("Fetching logs for all pods in namespace: %s\n", namespace)

	// 1. We first need to get the list of pods so we know their names and containers
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list pods for log extraction: %w", err)
	}

	logsDir := filepath.Join(outputDir, "logs")
	os.MkdirAll(logsDir, 0755)

	// 2. Loop over every Pod
	for _, pod := range podList.Items {

		// 3. Combine both standard containers and init containers into one list to check
		allContainers := append(pod.Spec.InitContainers, pod.Spec.Containers...)

		// 4. Loop over every container inside this specific Pod
		for _, container := range allContainers {

			// Define a specific configuration options struct for the log request
			logOptions := &corev1.PodLogOptions{
				Container: container.Name,
			}

			// 5. Create the API Request (This doesn't execute the request yet)
			req := clientset.CoreV1().Pods(namespace).GetLogs(pod.Name, logOptions)

			// 6. Stream the connection open
			podLogsStream, err := req.Stream(context.Background())
			if err != nil {
				fmt.Printf("Warning: Could not open log stream for pod %s/%s: %v\n", pod.Name, container.Name, err)
				continue
			}

			// 7. Read everything from the stream into memory and close it immediately
			logBytes, err := io.ReadAll(podLogsStream)
			podLogsStream.Close() // Crucial to prevent memory leaks!
			if err != nil {
				continue
			}

			// 8. Save the file using a clear naming convention: podname-containername.log
			logFileName := fmt.Sprintf("%s-%s.log", pod.Name, container.Name)
			filePath := filepath.Join(logsDir, logFileName)
			_ = os.WriteFile(filePath, logBytes, 0644)
		}
	}

	return nil
}
