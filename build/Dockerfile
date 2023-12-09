FROM golang:1.21-alpine as builder

WORKDIR /app

COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download

COPY "./pkg" "./pkg"
COPY "./internal/app" "./internal/app"
COPY "./internal/config" "./internal/config"
COPY "./internal/product" "./internal/product"
COPY "./internal/warehouse" "./internal/warehouse"
COPY "./cmd" "./cmd"

RUN go build -o ./cmd ./cmd

CMD ["./cmd/cmd"]