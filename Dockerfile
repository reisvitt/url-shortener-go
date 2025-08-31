# build step
FROM golang:1.22.6-bookworm AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /run-app .

# final step
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /run-app /usr/local/bin/
EXPOSE 8080
CMD ["/usr/local/bin/run-app"]
