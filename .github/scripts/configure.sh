#!/bin/bash

# Exit on error
set -e

aws eks --region us-east-2 update-kubeconfig --name CloudInfraQAEKS

pod=$(kubectl get pod -l app.kubernetes.io/name=alationfc -o jsonpath="{.items[0].metadata.name}")

echo "Pod name is: "$pod

kubectl wait --for=condition=ready pod $pod

alation_config_commands1="alation_conf -s True alation.ocf.acm.k8s && \
alation_conf alation.feature_flags.enable_alation_analytics_v2 -s True && \
alation_conf alation_analytics-v2.pgsql.config.host -s postgres && \
alation_conf alation_analytics-v2.rmq.config.host -s rabbitmq && \
alation_conf alation_analytics-v2.pgsql.config.port -s 5432 && \
alation_supervisor restart web:* celery:*"

alation_config_commands2="alation_conf alation_analytics-v2.pgsql.password -s 'password@123'"


command1="chroot /opt/alation/alation /bin/sh -c 'su - alationadmin -c \""$alation_config_commands1"\"'"

command2="chroot /opt/alation/alation /bin/sh -c 'su - alation -c \""$alation_config_commands2"\"'"

kubectl exec $pod -- bash -c "$command2"
kubectl exec $pod -- bash -c "$command1"
