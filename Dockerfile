# Use the official Golang image to create a build stage.
FROM golang:1.21 AS build-stage

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go mod and sum files and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project and build it.
COPY . .
RUN go build -o /go-canteen-project

# Set the entry point for the container.
CMD ["/go-canteen-project"]
