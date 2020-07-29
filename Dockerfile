FROM golang:1.14 as builder

WORKDIR /build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix nocgo -o app-registry ./cmd/qrappstore

FROM alpine:latest
WORKDIR /app
RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*
COPY --from=builder /build/app-registry ./
RUN chmod +x ./app-registry
ENTRYPOINT ["./app-registry"]