# Use an official Go runtime as the base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the source code from the host to the container
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080 to the host
EXPOSE 8080

# Run the application
CMD ["./main"]
