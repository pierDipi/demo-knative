#!/usr/bin/env bash

kubectl create namespace kafka --dry-run -o yaml | kubectl apply -f -

sleep 5

echo "Applying Strimzi Cluster Operator file"
kubectl apply -n kafka -f $(dirname $0)/strimzi-cluster-operator.yaml

sleep 5

kubectl -n kafka apply -f $(dirname $0)/kafka.yaml
kubectl apply -n kafka -f $(dirname $0)/user-tls.yaml
kubectl apply -n kafka -f $(dirname $0)/user-sasl-scram-512.yaml
