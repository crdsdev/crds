# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.13 as builder

# Copy internal libraries.
COPY internal /internal

# Create and change to the service directory.
WORKDIR /services/static

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY ./services/static/go.* ./
RUN go mod download

# Copy local code to the container image.
COPY ./services/static .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o static

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /services/static ./
COPY ./services/static/public ./public

# Run the web service on container startup.
ENTRYPOINT ["/static"]