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

func main() {
	ctx := context.Background()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	server := newServerCtx()
	pb.RegisterSchedulerServer(s, server)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal("failed to serve %v", err)
		}
	}()

	log.Println("Server listening on port", port)

	node, err := server.addResource("./utils/resource.dat")
	if err != nil {
		log.Fatal(err)
	}

	root, err := server.dag.Get(ctx, node.Cid())
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range root.Links() {
		log.Println(l.Name)
		blk, _ := server.blocks.GetBlock(ctx, l.Cid)
		log.Println(blk.RawData())
		log.Println("-------")
	}

	select {}
}
