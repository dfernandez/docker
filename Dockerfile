FROM golang:1.4.2

MAINTAINER david fernandez

EXPOSE 80

ADD . /code

RUN go get github.com/garyburd/redigo/redis
RUN go build /code/ping_redis.go

ENTRYPOINT ./ping_redis
