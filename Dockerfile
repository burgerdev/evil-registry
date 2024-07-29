FROM golang:1.22 AS builder

COPY registry /src

WORKDIR /src

RUN CGO_ENABLED=0 go build -o /registry /src

FROM alpine:3.20

COPY --from=builder /registry /registry

ENTRYPOINT ["/registry"]
