FROM golang:1.18-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Add required packages
RUN apk add --update git curl bash

WORKDIR /app

COPY . .

RUN go build -o app

# Expose port
EXPOSE 9090

# Start app
CMD ["./app"]