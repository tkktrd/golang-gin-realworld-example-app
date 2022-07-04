## Modules caching
FROM golang:1.18.3-alpine3.16 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

## Builder
FROM golang:1.18.3-alpine3.16 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/app

## Server
FROM scratch
COPY --from=builder /bin/app /app
CMD ["/app"]