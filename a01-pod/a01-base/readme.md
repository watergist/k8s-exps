```shell
docker buildx build --push . -t watergist/k8s-exps:a01-pod-a01-base --build-arg APP_DIR="/a01-pod/a01-base"
helm3 upgrade --install charts
kns a01-pod-a01-base
```
