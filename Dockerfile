FROM golang:1.19-alpine
WORKDIR /app
ADD . .
RUN go build -o app ./cmd/cloudrun/.
EXPOSE 8080
CMD ["/app/app"]