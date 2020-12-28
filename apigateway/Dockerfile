FROM golang:1.15.2-alpine3.12 as debug
RUN apk --update add openssh-client
RUN apk --update add openssl
WORKDIR /app/cmd
COPY ./cmd/install-terraform.sh .
COPY ./cmd/install-ansible.sh .
RUN ./install-terraform.sh
RUN ./install-ansible.sh
ENV CGO_ENABLED 0
WORKDIR /app/
COPY . .
RUN go build -gcflags "all=-N -l" -o apigateway ./cmd 
RUN go get github.com/go-delve/delve/cmd/dlv
EXPOSE 3001 2345
CMD ["dlv", "--listen=:2345", "--headless=true",  "--api-version=2","--log=true","--log-output=debugger,debuglineerr,gdbwire,lldbout,rpc", "--accept-multiclient=true", "debug", "/app/cmd"]

FROM golang:1.15.2-alpine3.12 as dev
RUN apk --update add openssh-client
RUN apk --update add openssl
WORKDIR /app/cmd
COPY ./cmd/install-terraform.sh .
COPY ./cmd/install-ansible.sh .
RUN ./install-terraform.sh
RUN ./install-ansible.sh
ENV CGO_ENABLED 0
WORKDIR /app/
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
EXPOSE 3001
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait
CMD /wait && CompileDaemon -directory=cmd/. -log-prefix=false -build="go build -o apigateway" -command="./cmd/apigateway"

FROM golang:1.15.2 as builder
RUN apt update && apt install git ca-certificates && update-ca-certificates
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