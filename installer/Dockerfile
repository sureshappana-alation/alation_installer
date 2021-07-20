# This docker file can be used to build the installer binary for linux using this command "docker build -o out ."
FROM golang:1.16.4-buster AS build
WORKDIR /src
COPY src/ .
COPY go.mod go.mod
COPY go.sum go.sum
RUN go build -o /out/installer *.go
COPY --from=compress /installer.tar.gz /
