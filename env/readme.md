# KIND cluster
```cd env/kind```
* Install KIND cluster deployment tool
    ```shell
    bash -ex install.sh
    ```
  
* Create a cluster
  ```shell
  kind create cluster --name k8s-exps
  # KUBECONFIG will be updated by itself
  ```
  
* Enable IP for loadBalancers (works only with KIND)
  ```
  IP_SET_NUMBER=3 CLUSTER_NAME=k8s-exps bash -ex metallb.sh
  ```