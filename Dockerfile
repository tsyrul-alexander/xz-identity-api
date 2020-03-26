FROM golang:latest
RUN mkdir -p /go/src/github.com/tsyrul-alexander/identity-web-api/
WORKDIR /go/src/github.com/tsyrul-alexander/identity-web-api/
COPY . /go/src/github.com/tsyrul-alexander/identity-web-api/
RUN go get github.com/tools/godep
RUN godep restore
EXPOSE 8080
ENTRYPOINT go run main.go