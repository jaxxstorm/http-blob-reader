FROM golang:alpine as build
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/app

FROM scratch
ENV GIN_MODE=release
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY http-blob-reader /app/http-blob-reader

ENTRYPOINT [ "/app/http-blob-reader" ]
