#prod config
FROM golang:1.15.6-alpine3.12 as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN mkdir /build
COPY . /build/
WORKDIR /build/cmd

RUN go mod download
RUN go mod verify
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o apigateway

FROM alpine 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/cmd/apigateway /app/
COPY --from=builder /build/cmd/install-terraform.sh /app/
COPY --from=builder /build/cmd/install-ansible.sh /app/
RUN apk --update add openssh-client

WORKDIR /app/

RUN chmod +x install-terraform.sh
RUN chmod +x install-ansible.sh

RUN /app/install-terraform.sh
RUN /app/install-ansible.sh

CMD [ "./apigateway" ]