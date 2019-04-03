package main

import (
	"context"
	pb "github.com/hashmatter/pstashio/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	serverAddr = "localhost:8080"
)

type Peer struct {
	client     pb.SchedulerClient
	conn       *grpc.ClientConn
	serverAddr string
}

func NewPeer(sAddr string) (*Peer, error) {
	conn, err := grpc.Dial(sAddr, grpc.WithInsecure())
	if err != nil {
		return &Peer{}, err
	}

	cli := pb.NewSchedulerClient(conn)

	p := Peer{
		client:     cli,
		conn:       conn,
		serverAddr: sAddr,
	}

	return &p, nil
}

func main() {
	peer, err := NewPeer(serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer peer.conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.RegisterPeerRequest{}
	res, err := peer.client.RegisterPeer(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
