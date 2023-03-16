FROM golang:1.17

WORKDIR /app

RUN \
    apt-get update && apt-get upgrade -y &&\
    apt-get -y install postgresql-client

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN chmod a+x waitPostgres.sh

#RUN pwd
RUN ls

RUN go build -o ./bin/tg_weather_bot ./cmd/app/app.go

EXPOSE 8080