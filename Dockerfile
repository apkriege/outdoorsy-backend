# Use a base Go image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules definition and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Set the entrypoint command
CMD ["./main"]