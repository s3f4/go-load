#prod config
FROM golang:1.15.6-alpine3.12 as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN mkdir /build
COPY . /build/
WORKDIR /build/cmd

RUN go mod download
RUN go mod verify
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o eventhandler

FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/cmd/eventhandler /app/
WORKDIR /app/
CMD [ "./eventhandler" ]
