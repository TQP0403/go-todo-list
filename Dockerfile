# syntax=docker/dockerfile:1
FROM ubuntu:22.04
FROM golang:1.21.1-alpine

# Download Go modules
COPY go.mod go.sum ./

RUN go mod download

RUN apk update && apk add --no-cache 'git=~2'

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/main .

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

# Run
CMD ["/go/main"]