#!/usr/bin/env bash

kubectl apply -f $(dirname $0)/cass-operator-manifests-v1.19.yaml
kubectl apply -f $(dirname $0)/cluster.yaml

kubectl wait --for=condition=Ready pods --all -n cass-operator

CASS_PASSWORD=$(kubectl get secret --template='{{.data.password}}' -n cass-operator cluster-superuser | base64 -d)
CASS_USERNAME=$(kubectl get secret --template='{{.data.username}}' -n cass-operator cluster-superuser | base64 -d)

kubectl create secret -n users-service generic cassandra-cluster-superuser \
  --from-literal=username="${CASS_USERNAME}" \
  --from-literal=password="${CASS_PASSWORD}"

kubectl create secret -n kafka generic cassandra-cluster-superuser \
  --from-literal=username="${CASS_USERNAME}" \
  --from-literal=password="${CASS_PASSWORD}"
