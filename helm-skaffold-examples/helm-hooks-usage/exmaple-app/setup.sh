#!/bin/bash

# Generate Pem data
pemdata="random-pem-data"

# Example to patch existing secret
# patch="[{\"op\":\"replace\",\"path\":\"/data/pem\",\"value\":\"$pemdata\"}]"
# kubectl patch secret secret1-alation-helm-getting-started-alation-helm-chart --type=json -p="$patch"

# To generate new secret
echo $1
kubectl create secret generic $1 --from-literal=pem=$pemdata
