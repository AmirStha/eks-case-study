#!/bin/bash

sudo yum install nano -y

curl --silent --location "https://github.com/weaveworks/eksctl/releases/download/v0.147.0/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin
eksctl version

export VERIFY_CHECKSUM=false
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash