## Run
```
go run main.go
```

## Build Docker
docker build -t docker_golang .
## Run Docker
docker run --name dockerPlaygolang -d -p 8080:8080 docker_golang
## Access docker image
 docker exec -it docker_golang bash
