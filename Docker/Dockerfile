# FROM golang:1.18.2-bullseye
FROM golang:1.18.2-alpine3.15


RUN mkdir /ServerDirectory

ADD . /ServerDirectory

WORKDIR /ServerDirectory

RUN go build  service.go

CMD ["/ServerDirectory/service"]

# type : docker build -t golangserver .
# to build container

# type "docker images"  to show containers list

# type "docker run -p "9090:8080" -it golangserver-alpine" to run your container on the docker