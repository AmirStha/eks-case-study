kubectl apply -f web-server/templates/html-configmap.yaml
kubectl apply -f web-server/templates/nginx-configmap.yaml
kubectl apply -f web-server/templates/nginx-deployment.yaml
kubectl apply -f web-server/templates/nginx-service.yaml
kubectl apply -f web-server/templates/nginx-ingress.yaml