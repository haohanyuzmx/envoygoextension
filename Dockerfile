FROM golang:1.21
WORKDIR /source
COPY ./go.mod ./
RUN go mod download
ENTRYPOINT go build -o simple.so -buildmode=c-shared .