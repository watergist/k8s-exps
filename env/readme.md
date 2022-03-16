Setup your environment as follows

```shell
cd env/a00-kind
bash -ex install.sh
kind create cluster --name isto
## Mess with your kubeConfig here
IP_SET_NUMBER=1 CLUSTER_NAME=isto bash -ex metallb.sh
```