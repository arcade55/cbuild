FROM alpine

COPY gopath/bin/xxx /go/bin/build

ENTRYPOINT /go/bin/build
