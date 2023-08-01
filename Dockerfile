# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files to the container
COPY . .

# Build the Go program (assuming the main.go is in /cmd/main.go)
RUN go build ./cmd/main.go

# Set the entry point to start the built binary
ENTRYPOINT ["./main"]

# Copy the root/wallet directory from the host system into the image
COPY /wallet /app/wallet