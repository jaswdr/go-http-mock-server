FROM golang:alpine

LABEL maintainer "Jonathan A. Schweder <jonathanschweder@gmail.com>"

COPY ./gomock.go /go/src

WORKDIR /go/src

RUN go build -o /usr/bin/gomock gomock.go \
    && chmod +x /usr/bin/gomock

ENTRYPOINT ["gomock"]
