# Rest test

## Running the app

To run this exmaple, from the root of this project:

```sh
go get github.com/gorilla/mux
go run *.go
```

## Build app for alpine docker container

```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rest-api-test .
```
