FROM golang:alpine as builder

LABEL authors="Christian Muehlhaeuser: buluma@gmail.com"

WORKDIR /go/markscribe
COPY . .
RUN go build

FROM alpine

COPY --from=builder /go/markscribe/markscribe /go/bin/markscribe

ENTRYPOINT ["/go/bin/markscribe"]
