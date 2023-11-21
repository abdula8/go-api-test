# Use a lightweight base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
# COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 9090

# Command to run the executable
CMD ["./main"]