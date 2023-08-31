# Stage 1: Build the Go binary
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o app

# Stage 2: Create a minimal image for the final app
FROM alpine:latest

WORKDIR /app

RUN apk update && apk add subversion --no-cache
RUN mkdir files

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]






# # Use an official Go runtime as the base image
# FROM golang:1.19-alpine

# # Set the working directory inside the container
# WORKDIR /app

# RUN apk update && apk add subversion

# # Copy the Go application source code into the container
# COPY . .

# # Build the Go application
# RUN go build -o app

# # Expose the port the application runs on
# EXPOSE 8080

# # Command to run the application
# CMD ["./app"]
