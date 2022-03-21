# Experiments
Important and common features of kubernetes apis, all tested with their own dedicated customized application and environments.
The below experimentation method might not come in handy at first,
but it ensures the isolation of different environments and manageability of numerous experiments with minimum code. 

```shell
# Different/Same experiments needs the version to be different to co-exist
VERSION=1
# Manually set exp name to be used throughout the build and deployment
# This is the directory name for the experiment
EXP_NAME=""

docker buildx build --push . -t watergist/k8s-exps:exp-"$EXP_NAME" --build-arg APP_DIR=exp"/$EXP_NAME"

kubectl create ns $EXP_NAME
kns $EXP_NAME
helm3 upgrade --install --set version="$VERSION" --set expName="$EXP_NAME" --set subset="v1" "$EXP_NAME-$VERSION" exp/$EXP_NAME/charts --namespace default

# Uninstall all the versions of current experiment 
helm3 list --namespace default -q | grep $EXP_NAME | xargs helm3 uninstall --namespace default
```