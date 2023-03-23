[ -z $CLUSTER_NAME ] && echo Need a cluster name && exit 1
[ -z $IP_SET_NUMBER ] && echo Need a Set-Number for ip range && exit 1
# get cidr range
CIDR=$(docker network inspect kind | jq ".[].Containers | .[] | select(.Name==\"$CLUSTER_NAME-control-plane\") | .IPv4Address " | grep -o -E "[\.0-9\/]*")
# all ips in the cidr range
# get last n+1 lines, in which last line is some arbitrary message, and rest n are ips
# at one time, get only 100 IPs
nmap -sL -n $CIDR | tail -"$(expr $IP_SET_NUMBER "*" 100 + 1 )" | head -100 | grep -o -E "[\.0-9]*" > ip-available-in-docker-network

# a string having an ip range
IP_RANGE="$(head -1 ip-available-in-docker-network )-$(tail -1 ip-available-in-docker-network )"

kubectl delete ns metallb-system || true
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml
kubectl label ns metallb-system ip-set-number=$IP_SET_NUMBER ip-range=$IP_RANGE
kubectl wait --namespace metallb-system \
                --for=condition=ready pod \
                --selector=app=metallb \
                --timeout=90s

curl https://kind.sigs.k8s.io/examples/loadbalancer/metallb-config.yaml > tempk8.yaml

## i would like to prefer yq next time.
#https://stackoverflow.com/questions/22960387/what-does-the-comma-in-sed-commands-stand-for
sed -ie  "/^\s\saddresses/,/---/ s/- .*/- $IP_RANGE/" tempk8.yaml
kubectl apply -f tempk8.yaml
rm tempk8.yaml ip-available-in-docker-network