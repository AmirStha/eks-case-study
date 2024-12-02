## Overview

This project demonstrates the deployment of a Kubernetes-based solution that includes the following components:

1. An EKS cluster
2. A MySQL database cluster with persistent storage.
3. A customized web server (Nginx) configured to dynamically display the serving pod’s details.
4. A Golang-based monitoring application that tracks pod lifecycle events.

The MySQL database cluster and the web server components are deployed and managed using Helm charts. The Golang application is deployed locally.
Note: The solution was implemented in an EKS (Elastic Kubernetes Service) cluster provisioned using A Cloud Guru's Sandbox environment ( The sandbox accounts are timeboxed and for the presentation, all components need to be reprovisioned ).

## Components and Implementation

### 1. EKS Cluster Setup

The Kubernetes cluster is provisioned on AWS using eksctl with the configuration file located in `infrastructure/eksctl/cluster.yaml`.

**Features:**
- Managed by eksctl, allowing rapid deployment of EKS clusters.
- Includes setup for AWS Load Balancer Controller to handle Kubernetes Ingress resources.
- IAM policies (iam-policy.yaml) are defined as a means of secure integration with AWS resources(Load Balancing).

**How to Deploy:**

Run the following script:
```
./scripts/01-init-cluster.sh
```

**Relevant screenshots located at `screenshots/eks`**

### 2. Database Deployment

A MySQL database is deployed with persistent storage using the Helm chart in `infrastructure/database/`.

**Features:**
- Persistent Storage: Ensured via PersistentVolume and PersistentVolumeClaim definitions.
- Secrets: Credentials for database access are securely stored in Kubernetes Secrets.
- Includes a test script to verify connectivity (`scripts/test-database.sh`).

**How to Deploy:**

```
./scripts/02-init-database.sh
```

Note: A helm chart can be found at `infrastructure/database/`

**Relevant screenshots located at `screenshots/database`**

### 3. Web Server Deployment

The Nginx web server is deployed using the Helm chart in `web-server/`. This multi-replica setup dynamically configures each pod to display its hostname.

**Features:**
- An `initContainer` modifies the HTML file at runtime to display the pod’s name.
- Ingress Configuration: Allows external access via AWS Load Balancer Controller.
- Custom Nginx Configuration: Managed via ConfigMaps.

**How to Deploy:**

```
./scripts/04-init-webserver.sh
```

Note: A helm chart can be found at `web-server/`

**Accessing the Web Server:**

- Retrieve the public URL:
```
kubectl get ingress
```
- Open the URL in a browser to view the webpage

**Relevant screenshots located at `screenshots/web-server`**

### 4. Monitoring Application

The monitoring-app/ directory contains a Golang-based monitoring application that tracks pod lifecycle events in the cluster.

**Features:**
- Event Logging: Logs creation, deletion, and updates of pods in real-time.
- Kubernetes API: Uses the client-go library to interact with the Kubernetes API server.

**How to Run locally:**

- Configure EKS config:
```
aws eks --region us-east-1 update-kubeconfig --name eks-acg  --profile acg
```

- Navigate to the monitoring-app via 
```
cd monitoring-app
```

- Run 
```
go run main.go
```

**Relevant screenshots located at `screenshots/monitoring-app`**

## Future Enhancements

1. Security

- Use separate namespaces for the database and the web server
- Implement network policies to control traffic between pods.
- Use AWS IAM Roles for Service Accounts (IRSA) for secure access to AWS resources such as AWS Load Balancer.
- A better way to store database credentials

2. Scaling
- Enable Horizontal Pod Autoscaler (HPA) for dynamic scaling of web server pods based on CPU or memory utilization.
- Add support for database clustering with read replicas.

3. Observability
- Integrate Prometheus and Grafana for metrics visualization.
- Centralize logging with tools like Fluentd or Loki.

4. Automation
- Use CI/CD tools to automate Helm chart deployments.
- Include automated rollback mechanisms in case of failed deployments.


