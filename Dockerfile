# Use the official Go image as a parent image
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy local code to the container image
COPY . ./

COPY go.mod go.sum ./

RUN go mod tidy && \
    go mod download

RUN CGO_ENABLED=0 go build -v -o webhook-server || (echo "Build failed" && exit 1)

# Use a minimal image for running the binary
FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /app/webhook-server /webhook-server

# Run the webhook server when the container starts
ENTRYPOINT ["/webhook-server"]
