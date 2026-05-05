FROM golang:alpine AS builder

LABEL authors="Michael Buluma: bulumaknight@gmail.com"

WORKDIR /go/markscribe
COPY . .
RUN CGO_ENABLED=0 go build -o markscribe .

FROM alpine

COPY --from=builder /go/markscribe/markscribe /usr/local/bin/markscribe

ENTRYPOINT ["/usr/local/bin/markscribe"]
