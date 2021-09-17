# > docker build -t playground .
FROM golang:alpine AS build-env

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o billboard-playground .

FROM alpine

RUN apk add --update ca-certificates
WORKDIR /app

COPY --from=build-env /build/.env .
COPY --from=build-env /build/billboard-playground .

EXPOSE 5000

ENTRYPOINT ["./billboard-playground"]
