FROM golang:latest
RUN mkdir /go/src/mytraining_backend
WORKDIR /go/src/mytraining_backend
COPY . .
CMD go run main.go 
