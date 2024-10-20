FROM golang:1.23-alpine AS builder

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/serve cmd/serve/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=builder --chown=nonroot:nonroot /app/bin/serve /app/serve

ENTRYPOINT ["/app/serve"]