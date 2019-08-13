### Build the code
FROM golang:1.12-alpine
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64
RUN apk add --update git
WORKDIR /go/src/git.sr.ht/~hokiegeek/iqhoarder
COPY . .
RUN go install -v -ldflags="-w -s" ./...

### Package it up
FROM alpine
EXPOSE 8078
COPY --from=0 /go/bin/iq-hoarder /
ONBUILD ENTRYPOINT ["/iq-hoarder"]
ENTRYPOINT ["/iq-hoarder"]