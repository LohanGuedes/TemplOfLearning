FROM golang:1.22-alpine as build

WORKDIR /app

COPY . .

RUN apk add --no-cache upx make
RUN go mod download && make production/deploy


FROM alpine:latest

WORKDIR /app

COPY --from=build /tmp/bin/api .
COPY --from=build /app .

EXPOSE 8080

ENTRYPOINT [ "./api" ]
