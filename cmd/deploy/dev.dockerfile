FROM golang:1.15-buster

RUN go get github.com/cespare/reflex

WORKDIR /app

COPY . .

ENTRYPOINT ["reflex", "-c", "./cmd/deploy/reflex.conf"]
