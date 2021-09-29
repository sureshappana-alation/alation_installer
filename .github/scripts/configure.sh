#!/bin/bash

# Exit on error
set -e

aws eks --region us-east-2 update-kubeconfig --name CloudInfraQAEKS

pod=$(kubectl get pod -l app.kubernetes.io/name=alationfc)

kubectl wait --for=condition=ready pod $pod

echo "Pod name is: "$pod

kubectl exec $pod -- bash -c "service alation shell && \
alation_conf -s True alation.ocf.acm.k8s && \
alation_conf alation.feature_flags.enable_alation_analytics_v2 -s True && \
alation_conf alation_analytics-v2.pgsql.config.host -s postgres && \
alation_conf alation_analytics-v2.rmq.config.host -s rabbitmq && \
alation_conf alation_analytics-v2.pgsql.config.port -s 5432 && \
"
