# STEP 1 build executable binary
FROM golang:1.11-alpine as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates
RUN adduser -D -g '' appuser

#Build API
COPY . $GOPATH/src/github.com/bosonic-code/api-mock/
WORKDIR $GOPATH/src/github.com/bosonic-code/api-mock/cmd/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/api 


USER appuser

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/ /go/bin/
COPY ./tls/server.crt /go/bin/
COPY ./tls/server.key /go/bin/
WORKDIR /go/bin 
ENTRYPOINT ["./api"]
