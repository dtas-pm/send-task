FROM golang:1.18

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

## install redis
#RUN apt-get update
#RUN apt-get -y install redis

# make wait-for-postgres.sh executable
RUN chmod +x ./cmd/wait-for-postgres.sh

## make wait-for-redis.sh executable
#RUN chmod +x ./cmd/wait-for-redis.sh

# build go app
RUN go mod download
RUN go build -o send-task ./cmd/main.go

CMD ["./send-task"]