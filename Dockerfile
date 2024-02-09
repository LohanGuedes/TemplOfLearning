FROM golang:1.22-alpine as build

WORKDIR /app

COPY . .

RUN apk add --no-cache \
    upx \
    curl \
    make \
    && go mod download \
    && make production/deploy

FROM scratch

WORKDIR /app

COPY --from=build /tmp/bin/api .
COPY --from=build /usr/bin/curl .

EXPOSE 8080

HEALTHCHECK --interval=10s --timeout=3s CMD ./curl -f http://localhost:8080/ || exit 1

ENTRYPOINT [ "./api" ]
