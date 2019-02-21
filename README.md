# Simple Go WASM

#### Prerequisites

| dep    | version   | notes             |
|:-------|:----------|:----------------- |
| Go     | >= 1.11.5 |                   |
| Chrome | >= 72     | Web Browser       |
| Python | >= 3.6    | Used as webserver |

#### build and run

1. Compile go to wasm

```
GOOS=js GOARCH=wasm go build -o main.wasm
```

2. run the webserver

> any webserver will do as long as the _application/wasm_ mimetype is available

```bash
$ python3.6 serve.py
serving at port 8080
localhost - - [15/Feb/2019 13:20:31] "GET / HTTP/1.1" 200 -
localhost - - [15/Feb/2019 13:20:32] "GET /wasm_exec.js HTTP/1.1" 200 -
localhost - - [15/Feb/2019 13:20:32] "GET /main.wasm HTTP/1.1" 200 -
localhost - - [15/Feb/2019 13:20:32] "GET /favicon.ico HTTP/1.1" 200 -
```

#### Notes

The `wasm_exec.js` was taken from [github.com/golang/go/misc/wasm](https://github.com/golang/go/tree/release-branch.go1.11/misc/wasm)
