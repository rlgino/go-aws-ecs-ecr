FROM golang:1.19-alpine as builder
# Application working directory ( Created if it doesn't exist )
WORKDIR /build
# Copy all files ignoring those specified in dockerignore
COPY . /build/
# Execute instructions on a new layer on top of current image. Run in shell.
RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main ./cmd/cloudrun

FROM alpine:3.9.4

# metadata for better organization
LABEL app="go-helloworld"
LABEL environment="production"
# Set workdir on current image
WORKDIR /app
# Leverage a separate non-root user for the application
RUN adduser -S -D -H -h /app appuser
# Change to a non-root user
USER appuser
# Add artifact from builder stage
COPY --from=builder /build/main /app/
# Expose port to host
EXPOSE 8080
# Run software with any arguments
ENTRYPOINT ["./main"]