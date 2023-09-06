# basic Go web server

`go run webserver.go`

## Dockerize

`go mod init test`

`go get`

`go mod tidy`

`docker build -t my-golang-server .` build the image

`docker image tag my-golang-server:latest my-golang-server:v1.0` command to create a new tag for our image

`docker run -p8088:8080 my-golang-server`

## Storage

To create a managed volume, run:

`docker volume create roach`

`docker volume list`

## Networking

creates a new bridge network named mynet

`docker network create -d bridge mynet`

`docker network list`

## Docker compose

When this command is run, Docker Compose would read the file docker-compose.yml, parse it into a data structure in memory, validate where possible, and print back the reconstruction of that configuration file from its internal representation

`docker compose config`

Let’s start our application and confirm that it is running properly.

`docker compose up --build`

`docker compose stop` stop the containers

`docker compose down` remove them

***INFO*** `docker compose` to see what other commands are available

## Docker Swarm

`docker system info`

looking for a message `Swarm: active`

If Swarm isn't running, simply type docker `swarm init`
