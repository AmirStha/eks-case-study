kubectl create secret generic mariadb-passwords \
    --from-literal=mariadb-root-password='amir123' \
    --from-literal=mariadb-password='amir123'

helm install mariadb \
    --set auth.database=ekshandson \
    --set auth.username=amirpj \
    --set auth.existingSecret=mariadb-passwords \
    oci://registry-1.docker.io/bitnamicharts/mariadb


helm install mariadb \
    --set auth.database=ekshandson \
    --set auth.username=amirpj \
    --set auth.existingSecret=mariadb-passwords \
    --set resources.limits.cpu=500m \
    --set resources.limits.memory=512Mi \
    --set resources.requests.cpu=100m \
    --set resources.requests.memory=128Mi \
    oci://registry-1.docker.io/bitnamicharts/mariadb