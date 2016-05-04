FROM golang:alpine

ADD . /go/src/github.com/Finciero/tiresias

RUN go get github.com/rodrwan/cat-grpc
RUN go install github.com/Finciero/tiresias

ENTRYPOINT /go/bin/tiresias

EXPOSE 3000
