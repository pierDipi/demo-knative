#!/usr/bin/env bash

ISTIO_VERSION=${ISTIO_VERSION:-1.8.3}

curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${istio_version} TARGET_ARCH=x86_64 sh -

$(dirname $0)/"istio-${istio_version}"/bin/istioctl install \
				--set profile=default -y

kubectl patch -n istio-system svc istio-ingressgateway -p '{"spec":{"type": "NodePort"}}'

kubectl apply -f $(dirname $0)/istio_ingressgateway_svc.yaml
