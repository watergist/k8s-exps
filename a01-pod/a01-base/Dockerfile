# docker buildx build --push ../.. -f $PWD/Dockerfile -t watergist/k8s-exps:"a01-pod_a01-base" --build-arg APP_DIR="/a01-pod/a01-base"
FROM watergist/golang:17.7 as build-dependencies

WORKDIR /code
COPY go.mod go.sum ./
RUN go mod download

FROM build-dependencies as copyfiles
ARG APP_DIR
COPY $APP_DIR/ ./
RUN find ./ -type f \
    \! -name "*.go" \! -name "*.mod" \! -name "*.sum" \
    -delete

FROM build-dependencies as build-heimdallr
WORKDIR /code
COPY --from=copyfiles /code ./
RUN go build -o /app/exp app.go

FROM ubuntu:20.04
WORKDIR /wd
COPY --from=build-heimdallr /app/exp /app/exp

RUN chmod -R a+rw .
ENTRYPOINT ["/app/exp"]