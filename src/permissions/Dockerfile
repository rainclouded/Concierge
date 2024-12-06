FROM golang:1.23.2-alpine3.20

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./api ./api
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./tests ./tests

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

RUN chown -R appuser:appgroup /app
USER appuser

# Command to run the binary
CMD ["./server"]
