FROM golang:1.6

RUN mkdir -p /go/src/github.com/jessemillar
ADD . /go/src/github.com/jessemillar/byudzhet

WORKDIR /go/src/github.com/jessemillar/byudzhet
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/byudzhet"]

EXPOSE 8000
