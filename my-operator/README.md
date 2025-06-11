# Kubernetes AI Security Operator

## Overview

The Kubernetes AI Security Operator is an experimental project designed to leverage generative AI and other tools to enhance the security posture of Kubernetes clusters. It aims to provide intelligent analysis of manifests, runtime security anomaly detection, vulnerability scanning insights, network policy auditing, secrets auditing, and pod security assessments.

This operator introduces a set of Custom Resource Definitions (CRDs) to manage these security tasks and integrates with external tools like Large Language Models (LLMs), vulnerability scanners (e.g., Trivy), and telemetry systems (e.g., OpenTelemetry) to provide comprehensive security feedback and actions.

## Current Status

**Alpha - Initial Development & Scaffolding**

*   Basic operator scaffolding complete.
*   Initial Custom Resource Definitions (CRDs) for core features have been defined in Go structs:
    *   `AIManifestAnalysis`
    *   `AISecurityAnomaly`
    *   `AIVulnerabilityReport`
    *   `AINetworkPolicyAudit`
    *   `AISecretsAudit`
    *   `AIPodSecurityAssessment`
*   A basic reconciliation loop for `AIManifestAnalysis` has been implemented (logs input, no action yet).
*   Placeholder integration points for LLM, Trivy, and OpenTelemetry have been created.
*   A basic Helm chart for deployment has been created.

**Known Issues:**

*   **`controller-gen` / `make manifests` Failure:** There is a persistent issue with `controller-gen` (the tool used by `operator-sdk` to generate CRD YAMLs, RBAC rules, etc.) failing with a Go type-checking error (`panic: runtime error: invalid memory address or nil pointer dereference` in `go/types.(*StdSizes).Sizeof(0x0, ...)`). This has been observed across multiple versions of `controller-gen` (v0.10.0, v0.11.3, v0.12.0).
    *   **Impact:** CRD YAMLs (except for `AIManifestAnalysis`, which was manually created) and RBAC configurations are not being generated correctly by the build tools. The Helm chart uses a manually created CRD for `AIManifestAnalysis` and a potentially stale RBAC role copied from the Kustomize bases.
    *   **Workaround:** For `AIManifestAnalysis`, the CRD YAML was manually created. Other CRDs exist as Go structs but not as installable YAML.

## Building the Operator

### Prerequisites
* Go (version 1.22 or later recommended)
* Docker (or another container runtime compatible with `docker build`)
* Access to a Kubernetes cluster (e.g., Kind, Minikube, or a cloud provider)
* `kubectl` configured for your cluster
* Helm (v3 or later)

### Build Steps

1.  **Build the operator binary:**
    While `make build` is a common target, ensure your Go environment is set up.
    ```bash
    go build -o bin/manager cmd/main.go
    ```

2.  **Build the Docker image:**
    The operator image name and tag are typically defined in the `Makefile` (e.g., `IMG ?= controller:latest`).
    ```bash
    make docker-build IMG=<your-registry>/my-operator:latest
    # Or, if you have a specific image name in mind:
    # docker build -t <your-registry>/my-operator:latest .
    ```
    (Note: The default `IMG` in the Makefile is `controller:latest`. You'll likely want to tag it with a proper registry and version for deployment).

## Deploying the Operator

The operator can be deployed using the provided Helm chart.

1.  **Ensure CRDs are available:**
    The Helm chart includes the `AIManifestAnalysis` CRD. If deploying other CRDs in the future (once the `controller-gen` issue is resolved), they would need to be applied to the cluster first or included in the Helm chart.

2.  **Install the Helm chart:**
    Navigate to the project root (`my-operator`).
    ```bash
    helm install my-operator-release ./my-operator-chart --set image.repository=<your-registry>/my-operator --set image.tag=latest
    ```
    Replace `<your-registry>/my-operator` with the actual image repository you pushed to, and `latest` with your image tag.

    You can customize the installation by overriding values in `my-operator-chart/values.yaml` or using `--set` flags. For example, to specify the ServiceAccount name if it wasn't the default:
    ```bash
    helm install my-operator-release ./my-operator-chart \
      --set image.repository=<your-registry>/my-operator \
      --set image.tag=latest \
      --set serviceAccount.name=my-custom-sa \
      --namespace my-operator-ns --create-namespace
    ```

## Development

*   API definitions are in `api/v1alpha1/`.
*   Controller logic is in `internal/controller/`.
*   Shared internal libraries (placeholders for LLM, scanners, telemetry) are in `pkg/`.

### Regenerating Code/Manifests
Normally, you would run `make manifests` to update CRDs and RBAC rules, and `make generate` for deepcopy functions. However, due to the `controller-gen` issue mentioned above, `make manifests` is currently failing.

This project was initialized with Operator SDK. For more details on Operator SDK, visit the [Operator SDK website](https://sdk.operatorframework.io/).
