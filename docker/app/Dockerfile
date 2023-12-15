FROM golang:1.21

WORKDIR /go/src

RUN apt-get update \
    && apt-get install -y --no-install-recommends bash

COPY . /go/src

# パッケージのインストール
RUN go mod init go-app \
    && go get -u github.com/spf13/cobra@latest \
    && go install github.com/spf13/cobra-cli@latest