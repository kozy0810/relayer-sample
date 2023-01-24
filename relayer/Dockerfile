FROM golang:1.19.4-alpine as dev

ENV ROOT=/go/src
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY . .
RUN go mod download

RUN go build -o /main

EXPOSE 8080

CMD ["/main"]