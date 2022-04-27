# Go Docker Tutorial
[Tutorial link](https://tutorialedge.net/golang/go-docker-tutorial/)

[Video link](https://www.youtube.com/watch?time_continue=617&v=lIbdPrUpGz4&feature=emb_logo)

# Run with Golang
```
go run main.go
```
which will kick of a server on http://localhost:10002.

# Run with Docker
```

docker build -t ii .
docker run -d --name ii --mount source=img,target=//usr/local/go/src/img-input/img -p 10002:10002 ii

```
