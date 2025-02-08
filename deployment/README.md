# Deployment
Minikube is a great way to run Kubernetes locally for testing and development. Below is a detailed step-by-step guide 
to setting up Minikube, deploying an application using Helm, and testing the deployment.

## 1. Install Dependencies
Before starting, ensure you have the required dependencies installed:
### Install Minikube
```shell
brew install minikube
```
### Install kubectl
Install the Kubernetes CLI (kubectl), which is required to interact with Minikube.
```shell
brew install kubectl  
```
### Install Helm
Helm is a package manager for Kubernetes.
```shell
brew install helm 
```

## 2. Start Minikube
Start the Minikube cluster with a specified driver (we use podman):
```shell
minikube start --driver=podman --container-runtime=cri-o
```
- Verify that the cluster is running:
```shell
kubectl cluster-info
```
- Check Minikube nodes:
```shell
kubectl get nodes
```

## 3. Enable Minikube Add-ons
Minikube comes with useful add-ons for development:
- Enable Ingress Controller for routing:
```shell
minikube addons enable ingress
```
- Enable Metrics Server (for Horizontal Pod Autoscaler):
```shell
minikube addons enable metrics-server
```
- Enable Dashboard (for UI monitoring):
```shell
minikube dashboard
```
Open minikube dashboard
http://127.0.0.1:57244/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/#/workloads?namespace=default

## 4. Create a Helm Chart
```shell
helm create currency-service
cd currency-service
```
This will generate a directory structure like:
```shell
currency-service/
├── charts/
├── templates/
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── ingress.yaml
│   ├── configmap.yaml
│   ├── secrets.yaml
├── values.yaml
├── Chart.yaml
├── README.md
```
## 5. Deploy to Minikube
Load Your Local Image into Minikube.
Open podman desktop and push image to minikube.
![podman.png](podman.png)

or you can use the following cli command:
```shell
minikube image tag docker.io/library/mysql:9.2 mysql:9.2 &&  minikube image push mysql:9.2
```

Deploy Helm Chart
Run the following command to deploy the Helm chart:
```shell
helm upgrade --install currency-service  ./currency-service 
```
Run the following command to test the rendered yaml:
```shell
helm upgrade --install currency-service  ./currency-service --dry-run
```

Check if the pods are running:
```shell
kubectl get po
```
Verify the services:
```shell
kubectl get svc
```

## 6. Expose and Test Your App
Option 1: Port Forwarding
If your service is not exposed publicly:
```shell
kubectl port-forward svc/atp-web 8080:80
```
Now, you can access the app at http://localhost:8080.

Option 2: Use Minikube IP
Find the Minikube cluster IP:
```shell
minikube ip
```
If using Ingress, edit your /etc/hosts file (Linux/macOS):
```shell
echo "$(minikube ip) my-app.local" | sudo tee -a /etc/hosts
```

If you want to open and access the database outside of the minikube.
## NodePort
```shell
kubectl get svc

minikube service mysql-node --url
http://127.0.0.1:53838
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.

```

Open connection through port 53838
![connection.png](connection.png)

## 7. Execute a command in a pod
```shell
kubectl exec currency-service-7f58898b7d-lw4kk  -- /app/service.currency -processor=currencyRateIngestor 
```

## 8. Monitoring and Debugging
- Check logs:
```shell
kubectl logs -l app=currency-service 
```
- Get pod details:
```shell
kubectl describe pod <pod-name>
```
- View running resources:
```shell
kubectl get all
```

## 9. Cleanup
- To delete your deployment:
```shell
helm uninstall my-app
```
- To stop Minikube:
```shell
minikube stop
```
- To delete everything:
```shell
minikube delete
```

## Summary of Workflow
1. Install dependencies (kubectl, helm, docker, minikube).
2. Start Minikube (minikube start).
3. Enable necessary add-ons (ingress, metrics-server).
4. Build and push your Docker image inside Minikube.
5. Deploy using Helm (helm upgrade --install).
6. Access your app using port forwarding or Ingress.
7. Execute a command in a pod
8. Monitor and debug with kubectl logs and kubectl get all.
9. Clean up when done (helm uninstall, minikube stop).


## References
- https://siweheee.medium.com/deploy-your-programs-onto-minikube-with-docker-and-helm-a68097e8d545
- https://podman-desktop.io/docs/minikube/pushing-an-image-to-minikube
- https://kubernetes.io/docs/tasks/tools/install-kubectl-macos/
- https://minikube.sigs.k8s.io/docs/drivers/podman/
