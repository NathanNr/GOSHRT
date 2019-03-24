FROM golang:1.11

LABEL maintainer="Nathan"

WORKDIR $GOPATH/src/github.com/NathanNr/GOSHRT
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["GOSHRT"]
