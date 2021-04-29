#!/usr/bin/env bash

set -xe

ISTIO_VERSION=${ISTIO_VERSION:-"1.8.3"}

#curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${ISTIO_VERSION} TARGET_ARCH=x86_64 sh -

$(dirname $0)/"istio-${ISTIO_VERSION}"/bin/istioctl install \
				--set profile=default -y

echo "Applying istio-ingressgateway"
kubectl delete svc -n istio-system istio-ingressgateway # ClusterIPs are immutable
kubectl apply -f $(dirname $0)/istio_ingressgateway_svc.yaml
