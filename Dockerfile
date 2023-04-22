# Start with a Golang 1.20 base image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the app
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
