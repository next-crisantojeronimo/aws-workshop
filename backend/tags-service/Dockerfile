FROM golang:1.14.8-alpine3.12

WORKDIR /tags-service

COPY . .

RUN go build -o tags-service.sh .

EXPOSE 4000
ENTRYPOINT ./tags-service.sh

# docker build -t cjeronimomx/tags-service:1.0.0 .