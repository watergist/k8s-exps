```shell
VERSION=1
docker buildx build --push ../.. -t watergist/k8s-exps:exp --build-arg APP_DIR="exp"
kubectl create ns $EXP_NAME
helm3 upgrade --install --set version="$VERSION" --set expName="$EXP_NAME" "$EXP_NAME-$VERSION" charts --namespace default
helm3 list --namespace default -q | grep $EXP_NAME | xargs helm3 uninstall --namespace default
kns $EXP_NAME
```