FROM golang:latest
WORKDIR /app
COPY . .
RUN go build
ENTRYPOINT test-health-server