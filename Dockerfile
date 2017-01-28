FROM golang

LABEL maintainer "Jonathan A. Schweder <jonathanschweder@gmail.com>"

COPY ./main.go /go/src

WORKDIR /go/src

RUN go build -o /usr/bin/gomock main.go

ENTRYPOINT ["gomock"]
