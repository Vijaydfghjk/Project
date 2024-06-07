# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Debugging: List contents of /app/pkg directory
RUN ls -R /app/pkg

# Build
RUN go build -o /app/bin ./cmd/main


# Set the binary as the entrypoint
ENTRYPOINT ["/app/bin"]
