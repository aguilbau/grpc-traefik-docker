package main

import (
	"log"
	"io/ioutil"
	"google.golang.org/grpc/credentials"
	"crypto/x509"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName = "world"
)

func main() {
     	FrontendCert, err := ioutil.ReadFile("/frontend.cert")
	if err != nil {
		log.Fatalln(err)
	}

	// Create CertPool
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(FrontendCert)

	// Create credentials
	credsClient := credentials.NewClientTLSFromCert(roots, "")

	authority := os.Getenv("AUTHORITY")
	address := os.Getenv("ADDRESS")

	// Dial with specific Transport (with credentials)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(credsClient))
	if authority != "" {
		opts = append(opts, grpc.WithAuthority(authority))
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	name := "World"
	r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
	   log.Fatalln(err)
	}

	log.Println(r)
}
