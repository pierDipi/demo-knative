## Create KinD cluster

```shell
kind create cluster --config kind/kind.yaml
```

## Install all components

```bash
./bin/install
```

***Verify installation***

```bash
kubectl get pods -A
```
Example output:
```
NAMESPACE            NAME                                          READY   STATUS    RESTARTS   AGE
istio-system         istio-ingressgateway-7675c8f974-nt74g         1/1     Running   0          64m
istio-system         istiod-86f494f976-p7skj                       1/1     Running   0          64m
kafka                my-cluster-entity-operator-5fd974964d-sd8gx   3/3     Running   2          20m
kafka                my-cluster-kafka-0                            1/1     Running   0          20m
kafka                my-cluster-kafka-1                            1/1     Running   0          20m
kafka                my-cluster-kafka-2                            1/1     Running   0          20m
kafka                my-cluster-zookeeper-0                        1/1     Running   0          21m
kafka                my-cluster-zookeeper-1                        1/1     Running   0          21m
kafka                my-cluster-zookeeper-2                        1/1     Running   0          21m
kafka                strimzi-cluster-operator-5f855cc555-wwxcq     1/1     Running   0          22m
knative-eventing     eventing-controller-d666b4657-zk69w           1/1     Running   0          64m
knative-eventing     eventing-webhook-778b6b8cf4-knzrl             1/1     Running   0          64m
knative-eventing     kafka-broker-dispatcher-754c758df4-ggkpl      1/1     Running   0          30m
knative-eventing     kafka-broker-receiver-77656c96dc-x7m76        1/1     Running   0          30m
knative-eventing     kafka-controller-5d89b5b555-jtpd4             1/1     Running   0          30m
knative-eventing     kafka-sink-receiver-868cffcdc7-4hc6x          1/1     Running   0          30m
knative-eventing     kafka-webhook-eventing-545f4bfc98-2glww       1/1     Running   0          30m
knative-serving      activator-86956bbd6f-2hfv8                    1/1     Running   0          63m
knative-serving      autoscaler-54cbd576f6-8gdp4                   1/1     Running   0          63m
knative-serving      controller-79c9cccd6f-4s2jg                   1/1     Running   0          63m
knative-serving      istio-webhook-56748b47-mtwxr                  1/1     Running   0          48m
knative-serving      networking-istio-5db557d5c4-w6xwj             1/1     Running   0          48m
knative-serving      webhook-5fd484cf4-hhqc8                       1/1     Running   0          63m
```

## Create APIServerSource

```shell

```