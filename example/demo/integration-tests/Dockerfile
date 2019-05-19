FROM golang:1.11

COPY . $GOPATH/src/github.com/bosonic-code/scaffold/
WORKDIR $GOPATH/src/github.com/bosonic-code/scaffold

CMD go test -v ./integration-tests/...
