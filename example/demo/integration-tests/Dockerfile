FROM golang:1.11

COPY . $GOPATH/src/github.com/bosonic-code/mock-api/example/demo/
WORKDIR $GOPATH/src/github.com/bosonic-code/mock-api/example/demo/

CMD go test -v ./integration-tests/...
