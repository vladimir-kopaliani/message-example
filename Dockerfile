FROM alpine:3.11.2

WORKDIR /opt

ADD ./build ./

ENTRYPOINT ["./main"]
