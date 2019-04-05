package main

import (
	"context"
	pb "github.com/hashmatter/pstashio/pb"
	"log"
)

type server struct {
	// a server keeps the current network state
	state *netstate
}

func newServerCtx() *server {
	return &server{
		state: newState(),
	}
}

// register peers in the network
func (*server) RegisterPeer(ctx context.Context, req *pb.RegisterPeerRequest) (*pb.RegisterPeerResponse, error) {
	log.Println("Received peer request: ", req)
	return &pb.RegisterPeerResponse{Status: "OK"}, nil
}

func (*server) GetProviders(ctx context.Context, req *pb.GetProvidersRequest) (*pb.GetProvidersResponse, error) {
	log.Printf("GetProviders (%v)\n", req)

	// checks the state of the peer

	// caclulates which blocks to replicate and from whom to provide plausible
	// deniability

	// creates replications manifest based on resources

	return &pb.GetProvidersResponse{}, nil
}
