#!/bin/bash

# Exit on error
set -e

aws eks --region us-east-2 update-kubeconfig --name CloudInfraQAEKS

pod=kubectl get pod -l app.kubernetes.io/name=alationfc

kubectl wait --for=condition=ready pod $pod

kubectl exec $pod -- bash -c ""
