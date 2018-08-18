FROM golang:latest
RUN mkdir /docker_example
ADD ../ /docker_example 
WORKDIR /docker_example
EXPOSE 8080
RUN go build -o main .
CMD ["/docker_example/main"]
# CMD ["go", "run", "main.go"]