package main

import (
	"context"
	pb "github.com/hashmatter/pstashio/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

type server struct{}

func (*server) RegisterPeer(ctx context.Context, req *pb.RegisterPeerRequest) (*pb.RegisterPeerResponse, error) {
	log.Println("Received", req)
	return &pb.RegisterPeerResponse{}, nil
}

func (*server) GetProviders(ctx context.Context, req *pb.GetProvidersRequest) (*pb.GetProvidersResponse, error) {
	log.Println("Received %v", req)
	return &pb.GetProvidersResponse{}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterSchedulerServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal("failed to serve %v", err)
		}
	}()

	log.Println("Server listening on port", port)
	select {}
}
