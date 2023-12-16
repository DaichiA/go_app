FROM golang:1.21

WORKDIR /go/cli

RUN go mod init cli \
    && go get -u github.com/spf13/cobra@latest \
    && go install github.com/spf13/cobra-cli@latest