FROM golang:1.23.2-alpine3.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-permissions

EXPOSE 8080

CMD ["/docker-permissions"]
