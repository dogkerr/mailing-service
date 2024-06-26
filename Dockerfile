# Step 1: Modules caching
FROM golang:1.22.2-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.22.2-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o /bin/app /app/cmd/app

# Step 3: Final
FROM alpine:3.12
COPY --from=builder /app/.env .
COPY --from=builder /bin/app /bin/app
ENV RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
CMD ["/bin/app"]
