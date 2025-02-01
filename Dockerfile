FROM golang:1.23-alpine AS builder

WORKDIR /build
COPY . .

RUN go mod tidy
RUN go mod download

# Build the application
RUN go build -o main ./cmd/app

############# New stage #######
FROM alpine:latest

# Install required system dependencies
RUN apk add --no-cache git ca-certificates

WORKDIR /root/

# Copy prebuilt binary file from previous stage
COPY --from=builder /build/main main

# bake .env to image
COPY --from=builder /build/.env .env

EXPOSE 8080
CMD ["./main"]