FROM golang:1.22-alpine as build

WORKDIR /app

COPY . .

RUN apk add --no-cache \
    upx \
    curl \
    make \
    && go mod download \
    && make production/deploy

FROM alpine:latest

WORKDIR /app

COPY --from=build /tmp/bin/api .
COPY --from=build /app .

RUN apk add --no-cache curl

EXPOSE 8080

HEALTHCHECK --interval=60s --timeout=3s CMD ["/bin/sh", "-c", "curl -f http://localhost:8080/v1/health || exit 1"]

ENTRYPOINT [ "./api" ]

