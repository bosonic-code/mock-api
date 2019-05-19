# STEP 1 build executable binary
FROM golang:1.11-alpine as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

#Build API
COPY . $GOPATH/src/github.com/bosonic-code/mock-api/
WORKDIR $GOPATH/src/github.com/bosonic-code/mock-api/cmd/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/api 


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/ /go/bin/
WORKDIR /go/bin 
ENTRYPOINT ["./api"]
