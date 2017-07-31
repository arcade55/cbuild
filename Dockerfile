FROM alpine

COPY gopath/bin/cbuild /go/bin/cbuild

ENTRYPOINT /go/bin/cbuild
