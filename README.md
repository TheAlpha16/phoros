# phoros

phoros is a Kubernetes-native, lightweight file server written in Go. It is designed to be scalable and efficient, capable of serving files with redundancy and persistent storage. Initially developed for use in CTF environments, phoros can be easily decoupled and used as a standalone solution. It integrates with Kubernetes and supports both native file storage and S3-compatible object storage (S3 support coming in the next release).

## Features

- **Lightweight & Efficient:** Optimized for minimal resource usage while handling file storage and transfer.
- **Kubernetes-Native:** Built specifically for Kubernetes environments, with easy deployment and configuration.
- **Helm Charts:** For easy configuration and deployment.
- **Scalable:** Supports autoscaling based on CPU/memory utilization.
- **Storage Options:** Native or S3-compatible object storage.
- **Traefik Proxy Integration:** Traefik proxy support for TLS certificates and external access through Kubernetes Ingress.

## Architecture

phoros is designed to run within a Kubernetes cluster. By default, it exposes a file server at a container port `9049` and can be configured to use a Kubernetes service to expose the server internally or externally. For external access, phoros leverages a Traefik proxy for TLS termination and ingress management.

## Getting Started

### Prerequisites

- Kubernetes Cluster
- Helm (for deployment)
- (Optional) S3-compatible object store (future release)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/TheAlpha16/phoros.git
   cd phoros
   ```

2. Customize the configuration by modifying the `values.yaml` file or by passing parameters during          installation

3. Install phoros using Helm:
   ```bash
   helm install phoros ./helm
   ```

4. Access phoros:
   - If ingress is enabled and properly configured, access phoros using the defined host (e.g., `https://example.com/files`).
   - If ingress is disabled, use the internal Kubernetes service (default type: `ClusterIP`).

## Usage

### Storage Configuration

You can choose between native file storage and S3-compatible object storage by setting the `objectStore` variable in the Helm `values.yaml` file.

- **Native Storage**:
  ```yaml
  storage:
    objectStore: native
  ```

- **S3-Compatible Storage** (available in future releases):
  ```yaml
  storage:
    objectStore: s3
  ```

## Configuration

phoros uses a `values.yaml` file for configuration. Below is a summary of the key configurable options:

- **Namespace**: Specify the Kubernetes namespace where phoros will be installed.
- **Event Timings**: Configure event start and end times if running phoros in a timed environment (e.g., CTFs).
- **Storage**: Define the storage class, access mode, and size. Choose between `native` or `s3` for the object store.
- **Secrets**: Set secrets for session management and admin access.
- **Service Type**: Specify the Kubernetes service type (`ClusterIP`, `NodePort`, or `LoadBalancer`).
- **Ingress**: Enable or disable ingress and define the host and paths.
- **Autoscaling**: Configure autoscaling parameters based on CPU/memory utilization.

For detailed configuration, check the [Helm `values.yaml`](./helm/phoros/values.yaml) file.

### Example `values.yaml`

```yaml
replicaCount: 1
namespace: phoros

event:
  startTime: 1729191091
  endTime: 1729291091
  postEvent: false

storage:
  storageClass: do-block-storage
  accessMode: ReadWriteOnce
  size: 1Gi
  objectStore: native
  mountPath: /etc/phoros

secrets:
  sessionSecret: your-session-secret
  adminSecret: your-admin-secret

service:
  type: LoadBalancer
  containerPort: 9049
  servicePort: 80

ingress:
  enabled: true
  host: "example.com"
  paths:
    - /files
    - /ping

autoscaling:
  enabled: false
  minReplicas: 1
  maxReplicas: 10
  targetCPUUtilizationPercentage: 80
```

## Contributing

Feel free to contribute to phoros by opening issues or submitting pull requests.

## Roadmap

- Add S3-compatible object storage support.
- Add tls support

## License

This project is licensed under the GNU General Public License.
