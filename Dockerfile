FROM golang:latest
RUN mkdir -p /go/src/identity-web-api/
WORKDIR /go/src/identity-web-api/
COPY . /go/src/identity-web-api/
RUN go get github.com/tools/godep
RUN godep restore
EXPOSE 8080
ENTRYPOINT go run main.go app_setting.go

FROM postgres:latest
FROM migrate/migrate:latest
RUN migrate -path: /go/src/identity-web-api/migration