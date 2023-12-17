FROM golang:1.21

WORKDIR /go/graphql

RUN go mod init graphql \
    && printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go \
    && go mod tidy \
    && go get -d github.com/99designs/gqlgen@latest
    # && go run github.com/99designs/gqlgen init \
    # && go mod tidy

# CMD [ "go", "run", "server.go" ]