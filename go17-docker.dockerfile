# for go builder: docker buildx build --push --target gobuilder . -f go17-docker.dockerfile -t watergist/golang:17.7

FROM ubuntu:20.04 as gobuilder
RUN apt-get update && apt-get install --no-install-recommends -y wget ca-certificates &&\
      wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz -O /tmp/golang.tar.gz -o /dev/null &&\
      tar -xf /tmp/golang.tar.gz --directory /usr/local &&\
      /usr/local/go/bin/go version && \
      apt-get purge wget -y && \
      apt-get autoremove -y && \
      apt-get clean && \
      find /var/lib/apt/lists -type f -delete && \
      find /var/cache -type f -delete && \
      find /var/log -type f -delete && \
      rm -rf /tmp/*

ENV PATH=/usr/local/go/bin:$PATH \
    GONOSUMDB="" \
    GOPROXY=https://proxy.golang.org,direct
