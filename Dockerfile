FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/adservice ./rpc/rta

FROM alpine:3.20

WORKDIR /app
RUN adduser -D -H -u 10001 appuser

COPY --from=builder /out/adservice /app/adservice

EXPOSE 8888
USER appuser

ENTRYPOINT ["/app/adservice"]
