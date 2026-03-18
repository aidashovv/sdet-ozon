FROM golang:1.25.6-bookworm AS builder
WORKDIR /sdet-ozon
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /sdet-ozon/cmd/cbr-service/exe ./cmd/cbr-service/

FROM alpine:3.23
WORKDIR /sdet-ozon
COPY --from=builder /sdet-ozon/cmd/cbr-service/exe /sdet-ozon
CMD ["/sdet-ozon/exe"]
