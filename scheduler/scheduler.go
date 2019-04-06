package main

import (
	"context"
	pb "github.com/hashmatter/pstashio/pb"
	"log"
)

// scheduler parameters
const (
	// size of the blocks in kb
	blockSize = 256
	// number of continguous blocks a peer will be responsible to provide
	blocksPeer = 10
	// obcurity factor is the number of extra (i.e. not requested) resources the
	// peer will request (some) blocks from
	obsFactor = 4
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

	return &pb.GetProvidersResponse{}, nil
}
