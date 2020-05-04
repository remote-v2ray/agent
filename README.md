## run

```
APIEndpoint=http://127.0.0.1/api?node=uuid-long-long-long ./v2wss
```

## desc

This project has replaced `vmess` and `websocket` module

The module of `ws` has changed `hub.go` file, add active user hook

The module of `vmess` has changed user validator

## how to build

```sh
# download vendor
go mod download
# build from vendor
cd cmd/v2wss && go build -o v2wss
```

## how to run

```
cd simple-server && go run main.go
APIEndpoint="http://127.0.0.1:1323/api?node=7777777777" ./v2wss
```
