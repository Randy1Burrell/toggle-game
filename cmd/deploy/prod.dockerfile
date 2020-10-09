# build stage
FROM golang:1.13-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go build -v -o main ./cmd/api/

# final stage
FROM alpine:latest
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
ENTRYPOINT ./main
