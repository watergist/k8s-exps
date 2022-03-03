[ -z $CLUSTER_NAME ] && echo Need a cluster name && exit 1
# get cidr range
CIDR=$(docker network inspect kind | jq ".[].Containers | .[] | select(.Name==\"$CLUSTER_NAME-control-plane\") | .IPv4Address " | grep -o -E "[\.0-9\/]*")
# all ips in the cidr range
# get last 5 lines, in which last line is some arbitrary message, and rest 4 are ips
nmap -sL -n $CIDR | tail -50 | head -49 | grep -o -E "[\.0-9]*" | tac > ip-available-in-docker-network

# a string having an ip range
IP_RANGE="$(tail -1 ip-available-in-docker-network )-$(head -1 ip-available-in-docker-network )"

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/master/manifests/namespace.yaml
kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)"
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/master/manifests/metallb.yaml

curl https://kind.sigs.k8s.io/examples/loadbalancer/metallb-configmap.yaml > tempk8.yaml

## search from bottom for - .*, at that line replace - .*
## i would like to prefer yq next time.
#https://stackoverflow.com/questions/22960387/what-does-the-comma-in-sed-commands-stand-for
sed -i "$,/- .*/ s/- .*/- $IP_RANGE/" tempk8.yaml
kubectl apply -f tempk8.yaml
rm tempk8.yaml