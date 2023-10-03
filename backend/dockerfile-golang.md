# Dockerfile pada Golang
Berikut sample Dockerizing pada Go Apps.

```dockerfile
FROM golang:1.20.6-alpine3.18

WORKDIR /nama-project

COPY go.mod /nama-project

COPY . /nama-project 

RUN go build -o /nama-project/bin/nama-binary /nama-project/main.go 

EXPOSE 8080

CMD ["/nama-project/bin/nama-binary"]

```
