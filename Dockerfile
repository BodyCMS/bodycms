FROM golang:1.19.4

# Install dependencies
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080