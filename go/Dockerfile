# Use an official Golang runtime as a parent image
FROM golang:1.21.3

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o api api.go

# Set the entry point with the command line argument
# CMD ["./api"]
ENTRYPOINT [ "./api"]
