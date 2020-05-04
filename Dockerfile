FROM golang:1.14-alpine as Build
WORKDIR /app
COPY . /app/
RUN set -e \
  && cd cmd/v2wss \
  && go build -mod=vendor -o v2wss

FROM alpine
# for https api
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=Build /app/cmd/v2wss/v2wss /app/v2wss
CMD ["./v2wss"]
