# Build the App
FROM golang:1.18.2-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o /app/appset-test-plugin

# Build the Image
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/appset-test-plugin /app/appset-test-plugin

EXPOSE 8080

USER 1001

CMD ["/app/appset-test-plugin"]
