FROM watergist/golang:17.7 as build-dependencies

WORKDIR /code
COPY go.mod go.sum ./
RUN go mod download

FROM build-dependencies as copyfiles

ARG APP_DIR
COPY $APP_DIR/ /code/cmd/
COPY pkg/ /code/pkg/
RUN find ./ -type f \
    \! -name "*.go" \! -name "*.mod" \! -name "*.sum" \
    -delete

FROM build-dependencies as builder
WORKDIR /code
COPY --from=copyfiles /code ./
RUN go build -o /app/exp /code/cmd/app/main.go

FROM ubuntu:20.04
WORKDIR /wd
COPY --from=builder /app/exp /app/exp

RUN chmod -R a+rw .
ENTRYPOINT ["/app/exp"]