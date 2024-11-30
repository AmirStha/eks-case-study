kubectl apply -f infrastructure/database/mysql-secret.yaml
wait

kubectl apply -f infrastructure/database/mysql-storage.yaml

sleep 180 &
pid=$!
wait $pid

kubectl apply -f infrastructure/database/mysql-deployment.yaml
