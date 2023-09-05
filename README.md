# basic Go web server

`go run webserver.go`

## Dockerize

`go mod init test`

`go get`

`go mod tidy`

`docker build -t my-golang-server .`

`docker run -p8088:8080 my-golang-server`
