# k8scollector

`k8scollector` is a lightweight, zero-dependency, standalone CLI tool built in Go designed for support engineers and cluster administrators. It allows you to instantly gather critical diagnostic configurations, workloads, namespace events, and pod logs from any Kubernetes cluster into a single structured directory for fast troubleshooting.

---

## Features

Without requiring `kubectl` to be installed on the host machine, `k8scollector` can connect directly to the Kubernetes API to collect:
* **Cluster/Namespace Specs:** Namespace configuration YAML.
* **Workloads:** Deployments, DaemonSets, StatefulSets, and CronJobs YAMLs.
* **Storage Layer:** PersistentVolumeClaims (PVCs) and filtered, relevant PersistentVolumes (PVs).
* **Network Components:** Ingress configurations.
* **Cluster Timeline:** A human-readable, chronological `events.log` table.
* **Application Logs:** Deep extraction of logs across all containers (including Init Containers) inside every pod.

---

## Getting Started

### Prerequisites
`k8scollector` automatically discovers your active Kubernetes cluster context using your local kubeconfig file. Ensure your cluster context is active or provide the path explicitly.

### Quick Start (Running the Release Binary)

1. **Download the Binary:**
   Go to the **Releases** section of this GitHub repository and download the compiled binary matching your operating system (Linux, macOS, or Windows).

2. **Make it Executable (Linux/macOS only):**
   ```bash
   chmod +x k8scollector

### Run the Collector:
By default, the tool will target the default namespace using your default ~/.kube/config profile:
```bash
./k8scollector
```

Command Line Flags
You can customize the target namespace or specify a custom kubeconfig path using the following flags:
```bash
# Target a specific namespace using short or long flags
./k8scollector -n production
./k8scollector --namespace production

# Point to a custom kubeconfig configuration file path
./k8scollector -n staging --kubeconfig /path/to/custom/kubeconfig
```

### Output Directory Structure
Once running completes, the tool creates an organized diagnostic directory named k8scollector-dump/ in your current working directory:
```bash
k8scollector-dump/
├── namespace.yaml
├── events.log            <-- Formatted kubectl-style events timeline
├── deployments/
│   ├── api-service.yaml
│   └── frontend.yaml
├── pods/
│   └── api-service-79ff-xyz.yaml
├── logs/
│   ├── api-service-79ff-xyz-main-app.log
│   └── api-service-79ff-xyz-init-migration.log
├── daemonsets/
├── statefulsets/
├── cronjobs/
├── ingress/
├── persistentvolumeclaims/
└── persistenvolumes/
```
