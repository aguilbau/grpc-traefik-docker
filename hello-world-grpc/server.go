package main

import (
	"log"
	"net"
	"io/ioutil"
	"crypto/tls"
	"google.golang.org/grpc/credentials"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Read cert and key file
	BackendCert, err := ioutil.ReadFile("/backend.cert")
	if err != nil {
	        log.Fatalln(err)
	}
	BackendKey, err := ioutil.ReadFile("/backend.key")
	if err != nil {
	        log.Fatalln(err)
	}

	// Generate Certificate struct
	cert, err := tls.X509KeyPair(BackendCert, BackendKey)
	if err != nil {
	        log.Fatalln(err)
	}

	// Create credentials
	creds := credentials.NewServerTLSFromCert(&cert)

	// Use Credentials in gRPC server options
	serverOption := grpc.Creds(creds)

	var s *grpc.Server = grpc.NewServer(serverOption)
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	defer s.Stop()

	log.Println("listening on port ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
