kubectl apply -f web-server/html-configmap.yaml
kubectl apply -f web-server/nginx-configmap.yaml
kubectl apply -f web-server/nginx-deployment.yaml
kubectl apply -f web-server/nginx-service.yaml