FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/baileyjm02/test-health-server

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the package
RUN go build

# This container exposes port 80 to the outside world
EXPOSE 80

# Run the executable
CMD ["test-health-server"]