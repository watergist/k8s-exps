Setup your environment as follows

```shell
cd a00-env/a00-kind
bash -ex install.sh
kind create cluster --name isto
## Mess with your kubeConfig here
CLUSTER_NAME=isto bash -ex metallb.sh
```