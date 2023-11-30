FROM golang:1.21-alpine as builder

# Create a non-root user
RUN adduser --disabled-password --gecos "" ports-user

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd cmd
COPY app app
COPY internal internal

RUN chown -R ports-user /app
USER ports-user

RUN go build -o ./build/webapp ./cmd/webapp

FROM alpine:3.17 as webapp

WORKDIR /app
COPY --from=builder /app/build/webapp .

CMD ["./webapp"]

