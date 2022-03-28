# Below are the go-to steps for deploying the base applications
* give version as ```VERSION=1```  Increment this when deploying more than one of these applications.
* Build docker image from root of the repository.

# Add helm repo
```shell
helm3 repo add k8s-exps https://watergist.github.io/k8s-exps
helm3 repo update k8s-exps
helm3 repo list # k8s-exps will be listed
```

# HTTP & HTTPS server
<details>
<summary>Create K8s-Secret for TLS</summary>
<ul>

```shell
mkdir tls && cd tls

# generate key pair first for a CA
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=example Inc./CN=example.com' -keyout ca.key -out ca.crt
# generate private-key and a csr having public key for a domain
openssl req -out tls.csr -newkey rsa:2048 -nodes -keyout tls.key -subj "/CN=httpbin.example.com/O=httpbin organization"
# request above created CA to use the csr and generate a certificate signed by it
openssl x509 -req -sha256 -days 365 -CA ca.crt -CAkey ca.key -set_serial 0 -in tls.csr -out tls.crt

rm tls.csr ca.crt ca.key
kubectl create secret tls tls-secret --key=tls.key --cert=tls.crt
cd .. && rm -r tls
```

</ul>
</details>

```shell
docker buildx build --push . -t watergist/k8s-exps:base-http --build-arg APP_DIR="base/http"

helm3 upgrade --install --set HTTPPorts="8080~9090~7070",HTTPSPorts="8081~9091~7071",HTTPTargetPorts="3000~4000~5000",HTTPSTargetPorts="3001~4001~5001" --set tlsSecret="tls-secret" --set version=${VERSION:-1} httpv${VERSION:-1} k8s-exps/base-http --version 0.0.5

# to uninstall all the releases
helm3 list | grep base-http | awk '{print $1}' | xargs helm3 uninstall
```

# TCP & UDP server
```shell
docker buildx build --push . -t watergist/k8s-exps:base-l4 --build-arg APP_DIR="base/l4"

# for tcp
helm3 upgrade --install --set TCPTargetPort="3001",TCPPort="8080" --set version=${VERSION:-1} tcpv${VERSION:-1} k8s-exps/base-l4 --version 0.0.5
# for udp
helm3 upgrade --install --set UDPTargetPort="3001",UDPPort="8080" --set version=${VERSION:-1} udpv${VERSION:-1} k8s-exps/base-l4 --version 0.0.5

# to uninstall all the releases
helm3 list | grep base-l4 | awk '{print $1}' | xargs helm3 uninstall
```