FROM golang:latest AS build

MAINTAINER mail@maltewildt.de

WORKDIR /src
COPY . /src

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o podman-kube-secrets ./main.go

FROM scratch
COPY --from=build /src/podman-kube-secrets /podman-kube-secrets
ENTRYPOINT ["/podman-kube-secrets"]