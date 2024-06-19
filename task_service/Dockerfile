FROM golang:1.19.5-alpine3.16 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -a -installsuffix cgo -o /main cmd/api/main.go

EXPOSE 8196

FROM alpine:3
COPY --from=builder main /app/main
ENTRYPOINT ["/app/main"]
