# k8s-exps

It's all about experimenting with Kubernetes.

* /env helps in setting up the base cluster for testing.
* /base holds the base charts that are able to dynamically adjust themselves via versions, subsets, ports. Multiple versions and subversions can co-exist.
* /exp wants to give a dedicated test environments for several Kubernetes features.
* /pkg obviously holds the generic go packages used to serve http or to implement other requirements.
