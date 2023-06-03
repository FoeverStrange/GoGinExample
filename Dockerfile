FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/GoGinExample
COPY . $GOPATH/src/GoGinExample
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./GoGinExample"]