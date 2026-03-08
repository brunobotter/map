FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache wget

COPY go.mod go.sum ./
RUN go mod download

COPY app/. .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/map .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /out/map ./map 
EXPOSE 8080
ENTRYPOINT ["./map"]