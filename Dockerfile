FROM golang:1.19-alpine

WORKDIR /app

# Download go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Get all go files
COPY *.go ./

# Get .env file
COPY .env ./

# Build binary
RUN go build -o /git-badges-golang

EXPOSE 8080

# Run binary
CMD [ "/git-badges-golang" ]