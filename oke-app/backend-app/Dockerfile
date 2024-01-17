FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY http/*.go ./http/
COPY db/*.go ./db/
COPY repo/*.go ./repo/
COPY crud/*.go ./crud/

RUN CGO_ENABLED=0 GOOS=linux go build -o demo-app

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

FROM alpine AS runner
WORKDIR /app

COPY --from=builder /app/demo-app ./
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["/app/demo-app"]
