#!/bin/bash

# Exit on error
set -e

aws eks --region us-east-2 update-kubeconfig --name CloudInfraQAEKS

alationfc_pod_name=$(kubectl get pod -l app.kubernetes.io/name=alationfc -o jsonpath="{.items[0].metadata.name}")

echo "Waiting for Alation fat container pod "$alationfc_pod_name " to become ready"

kubectl wait --for=condition=ready pod $alationfc_pod_name


aa_postgres_pwd_alation_conf_cmd="alation_conf alation_analytics-v2.pgsql.password -s 'password@123'"

aa_ocf_alation_conf_cmd="alation_conf -s True alation.ocf.acm.k8s && \
alation_conf alation.feature_flags.enable_alation_analytics_v2 -s True && \
alation_conf alation_analytics-v2.pgsql.config.host -s postgres && \
alation_conf alation_analytics-v2.rmq.config.host -s rabbitmq && \
alation_conf alation_analytics-v2.pgsql.config.port -s 5432 && \
alation_supervisor restart web:* celery:*"

aa_postgres_password_master_cmd="chroot /opt/alation/alation /bin/sh -c 'su - alation -c \""$set_aa_postgres_password_cmd"\"'"
aa_ocf_alation_conf_master_cmd="chroot /opt/alation/alation /bin/sh -c 'su - alationadmin -c \""$aa_ocf_alation_conf_cmd"\"'"

kubectl exec $alationfc_pod_name -- bash -c "$aa_postgres_password_master_cmd"
kubectl exec $alationfc_pod_name -- bash -c "$aa_ocf_alation_conf_master_cmd"
