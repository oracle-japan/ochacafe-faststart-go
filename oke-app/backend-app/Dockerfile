FROM golang:1.21 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY http/*.go ./http/
COPY db/*.go ./db/
COPY repo/*.go ./repo/
COPY crud/*.go ./crud/

RUN CGO_ENABLED=0 GOOS=linux go build -o demo-app

FROM scratch AS runner
WORKDIR /app

COPY --from=builder /app/demo-app ./

EXPOSE 8080

CMD ["/app/demo-app"]
