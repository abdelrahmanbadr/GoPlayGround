FROM golang:latest
RUN mkdir /project
ADD . /project/
WORKDIR /project/
RUN go build -o main .
ENV PORT 8080
EXPOSE 8080
CMD ["/project/main"]