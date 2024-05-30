FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prom-demo .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/prom-demo .
ENTRYPOINT ["./prom-demo"]
CMD ["start"]
