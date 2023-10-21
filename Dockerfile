FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .

RUN go build -o demo ./cmd/demo.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/demo /build/demo

CMD ["./demo"]
