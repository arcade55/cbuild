FROM alpine

COPY gopath/bin/build /go/bin/build

ENTRYPOINT /go/bin/build
