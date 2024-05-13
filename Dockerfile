# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

ENV RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache the module dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main .

RUN ls -l

RUN go run main.go

# Expose the port your application listens on (if applicable)
# EXPOSE 8080

# # Set the entrypoint to run the compiled binary
# ENTRYPOINT ["./main"]