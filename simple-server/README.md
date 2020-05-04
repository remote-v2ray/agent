## Run

```sh
go run main.go
```

```sh
# test
curl -sSL -H 'v2wss-action: GetPbConfig' http://127.0.0.1:1323/api | v2ray -format=pb -config=stdin:
```
