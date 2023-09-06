# pod-mutator
This is pod mutator webhook example where each new pod will have its own label which is predefined. 


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
   kubectl apply -f deploy/mutatingwebhook.yaml
   ```

2. Create a secret to store TLS certificates:

   ```shell
   kubectl create secret tls webhook-server-cert --cert=path/to/webhook-server.crt --key=path/to/webhook-server.key
   ```

## Testing the Mutating Webhook Server

1. Deploy a sample pod to test the webhook server:

   ```shell
   kubectl apply -f deploy/test-pod.yaml
   ```

2. Check the logs of the webhook server pod for debugging:

   ```shell
   kubectl logs -f <pod_name>
   ```

## Cleaning Up

To clean up the resources created for testing:

1. Delete the test pod:

   ```shell
   kubectl delete -f deploy/test-pod.yaml
   ```

2. Delete the MutatingWebhookConfiguration:

   ```shell
   kubectl delete -f deploy/mutatingwebhook.yaml
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
