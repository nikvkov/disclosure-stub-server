FROM golang:latest
WORKDIR /go/src/github.com/nikvkov/disclosure-stub-server
COPY . .
RUN go mod init
RUN go build -o stub-server stub-server.go
RUN ls
CMD ./stub-server