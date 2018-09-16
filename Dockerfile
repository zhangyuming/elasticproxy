FROM  golang:1.10-alpine3.8  AS builder
RUN apk add git
RUN go get github.com/tidwall/gjson
RUN go get github.com/tidwall/sjson

ADD . /go/src/elasticproxy/
RUN cd src/elasticproxy && go build -o elasticproxy




From alpine:3.8
COPY --from=builder /go/src/elasticproxy/elasticproxy /usr/bin/elasticproxy
CMD ["/usr/bin/elasticproxy"]                          
