FROM golang:1.14-alpine AS builder
WORKDIR /app
COPY . /app
RUN go build
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/files-content-exporter

FROM scratch
WORKDIR /app
COPY --from=builder /app/files-content-exporter .
CMD ["/app/files-content-exporter"]
