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
<<<<<<< HEAD
ENTRYPOINT ["./app-registry"]
=======
ENTRYPOINT ["./app-registry"]
>>>>>>> 6adcb63e5d9be034af6027da8acc0f6d04248aa3
