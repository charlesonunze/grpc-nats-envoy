package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/charlesonunze/grpc-nats-envoy/user-service/internal/db"
	"github.com/charlesonunze/grpc-nats-envoy/user-service/internal/rpc"
	"github.com/charlesonunze/grpc-nats-envoy/user-service/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Connect to NATS
	opts := nats.Options{
		AllowReconnect: true,
		MaxReconnect:   5,
		ReconnectWait:  5 * time.Second,
		Timeout:        3 * time.Second,
		Url:            os.Getenv("NATS_URL"),
	}

	nc, err := opts.Connect()
	if err != nil {
		fmt.Printf("err => %v", err)
		log.Fatal(err)
	}
	defer nc.Close()
	fmt.Println("nats connected")

	db.ConnectDB()
	defer db.CloseDB()

	s := grpc.NewServer()
	pb.RegisterUserServiceRPCServer(s, rpc.New(nc))
	reflection.Register(s)

	address := "0.0.0.0:5050"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Backend Error %v", err)
	}

	fmt.Println("userService is listening on", address)

	s.Serve(lis)
}
