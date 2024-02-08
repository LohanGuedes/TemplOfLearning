FROM golang:1.22-alpine as build

WORKDIR /app

COPY . .

RUN apk add --no-cache upx make
RUN go mod download && make production/deploy

FROM scratch

WORKDIR /app

COPY --from=build /app/tmp/bin/ .

EXPOSE 8080

ENTRYPOINT [ "./app/api" ]
