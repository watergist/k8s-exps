```shell
VERSION=1
docker buildx build --push . -t watergist/k8s-exps:base-l4 --build-arg APP_DIR="base/l4"
helm3 upgrade --install --set version="$VERSION" --set expName="$EXP_NAME" "$EXP_NAME-$VERSION" charts --namespace default
helm3 list --namespace default -q | grep $EXP_NAME | xargs helm3 uninstall --namespace default
```