# Build stage
FROM golang:1.20-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o releasr main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/releasr .

CMD [ "/app/releasr" ]