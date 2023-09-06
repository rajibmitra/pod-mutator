## Pod Mutator

In Kubernetes, a Mutating Webhook is a resource that allows you to intercept requests to the Kubernetes API server and modify them before they are persisted or admitted into the cluster. This capability is powerful for enforcing policies and making dynamic changes to resources.

In this project, we are creating a Mutating Webhook that specifically targets pods. The webhook server will intercept requests to create or update pods and apply custom logic to modify the pod specifications before they are admitted to the Kubernetes cluster.

![image](https://github.com/rajibmitra/pod-mutator/assets/1690251/b8cbf643-696e-499d-8aac-0fda1ccf0eb0)


## GoLang Code

The GoLang code in this project serves as the foundation for the Mutating Webhook Server. Let's break down the key components of the code:

1. **Dockerfile**: The `Dockerfile` defines how the GoLang application is containerized. It starts by using the official Go image as a parent image and sets up the working directory inside the container.

2. **Building the Webhook Server**:
   - The code is copied into the container.
   - The project's Go module files (`go.mod` and `go.sum`) are copied to ensure dependencies are available.
   - The Go module dependencies are downloaded using `go mod tidy` and `go mod download`.
   - The code is compiled for a specific target platform (in this case, `linux/amd64`) using `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build`. The resulting binary is named `webhook-server`.

3. **Creating the Final Image**:
   - A minimal Alpine Linux image is used as the base image.
   - Certificates are installed to ensure secure communication.
   - The compiled `webhook-server` binary is copied from the builder image to the final image.

4. **Entrypoint**: The `ENTRYPOINT` instruction specifies the command that should be executed when the container starts. In this case, it runs the `webhook-server`.

## Building the Docker Image

To build the Docker image, we use the `docker build` command with the `--platform linux/amd64` flag to specify the target platform. This ensures that the image is built for the `linux/amd64` platform, which is common for many development environments.

## Testing Locally

For local testing, we recommend using a Kind (Kubernetes in Docker) cluster. You can create a Kind cluster and deploy the Mutating Webhook Server as described in the README. The webhook server will intercept and modify pod specifications according to your custom logic.

Testing is facilitated by deploying a sample pod that triggers the webhook server. The server's logs can be checked to verify its behavior.

In summary, this project creates a Mutating Webhook Server using GoLang that intercepts pod creation and update requests in Kubernetes, applies custom logic to modify pod specifications, and provides a secure and reusable solution for policy enforcement and dynamic configuration in a Kubernetes cluster.


```markdown
# Mutating Webhook Server for Kubernetes

This project demonstrates how to create a mutating webhook server for Kubernetes. The webhook server modifies pod specifications before they are admitted to the cluster. In this example, we use a Kind (Kubernetes in Docker) cluster for local testing.

## Prerequisites

Before you begin, make sure you have the following prerequisites installed on your local machine:

- [Docker](https://www.docker.com/)
- [Kind](https://kind.sigs.k8s.io/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

## Building the Webhook Server

To build the webhook server, follow these steps:

1. Clone this repository:

   ```shell
   git clone https://github.com/yourusername/pod-mutator-webhook.git
   cd pod-mutator-webhook
   ```

2. Build the Docker image:

   ```shell
   docker build --platform linux/amd64 -t pod-mutator:0.1 .
   ```

3. Load the Docker image into your Kind cluster:

   ```shell
   kind load docker-image pod-mutator:0.1
   ```

## Creating a Kind Cluster

If you don't have a Kind cluster set up, you can create one with the following command:

```shell
kind create cluster --config kind-config.yaml
```

Ensure that you have a `kind-config.yaml` file with appropriate cluster configuration.

## Deploying the Mutating Webhook Server

1. Deploy the MutatingWebhookConfiguration:

   ```shell
   kubectl apply -f webhook-config.yaml
   ```

2. Create a secret to store TLS certificates:

   ```shell
   kubectl create secret tls webhook-server-cert --cert=path/to/webhook-server.crt --key=path/to/webhook-server.key
   ```

## Testing the Mutating Webhook Server

1. Deploy a sample pod to test the webhook server:

   ```shell
   kubectl apply -f test-pod.yaml
   ```

2. Check the logs of the webhook server pod for debugging:

   ```shell
   kubectl logs -f <pod_name>
   ```

## Cleaning Up

To clean up the resources created for testing:

1. Delete the test pod:

   ```shell
   kubectl delete -f test-pod.yaml
   ```

2. Delete the MutatingWebhookConfiguration:

   ```shell
   kubectl delete -f webhook-config.yaml
   ```

3. Delete the secret:

   ```shell
   kubectl delete secret webhook-server-cert
   ```

4. If you no longer need the Kind cluster, you can delete it:

   ```shell
   kind delete cluster
   ```

## Contributing

Feel free to contribute to this project by opening issues or creating pull requests. Your feedback and contributions are welcome!

## License

This project is licensed under the Apache Licence, feel free to use, modify and share. 
```

Please replace placeholders like `yourusername`, `path/to/webhook-server.crt`, and `path/to/webhook-server.key` with your actual values and paths. Additionally, make sure you have the appropriate Kubernetes configuration files for your Kind cluster.
