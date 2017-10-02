FROM golang:1.9

RUN go get google.golang.org/grpc

COPY ./server.go /go/src/google.golang.org/grpc/examples/helloworld/greeter_server/main.go
COPY ./client.go /go/src/google.golang.org/grpc/examples/helloworld/greeter_client/main.go
RUN cd / && go build /go/src/google.golang.org/grpc/examples/helloworld/greeter_server/main.go && mv /main /server
RUN cd / && go build /go/src/google.golang.org/grpc/examples/helloworld/greeter_client/main.go && mv /main /client

EXPOSE 50051
CMD ["/server"]
