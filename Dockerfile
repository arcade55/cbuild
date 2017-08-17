FROM alpine

RUN apk --update add ca-certificates

COPY gopath/bin/cbuild /go/bin/cbuild

ENTRYPOINT /go/bin/cbuild
