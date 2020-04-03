FROM golang:latest
RUN mkdir -p /go/src/github.com/tsyrul-alexander/xz-identity-api/
WORKDIR /go/src/github.com/tsyrul-alexander/xz-identity-api/
COPY . /go/src/github.com/tsyrul-alexander/xz-identity-api/
RUN go get github.com/tools/godep
RUN godep restore
EXPOSE 8080
ENTRYPOINT go run main.go