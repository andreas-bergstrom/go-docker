FROM golang:1.20-buster AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o /hello_world

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /hello_world /hello_world

ENTRYPOINT ["/hello_world"]