FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o cloudflare-dynamic-dns

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/cloudflare-dynamic-dns app

ENTRYPOINT ["/app"]