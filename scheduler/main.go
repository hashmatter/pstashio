package main

import (
	"context"
	pb "github.com/hashmatter/pstashio/pb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
)

const (
	port           = ":8080"
	resources_path = "./utils/"
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

	// loads all the blocks into the server's blockstore asynchronously
	go func() {
		// get all files in resource path
		rs, _ := ioutil.ReadDir(resources_path)
		for _, r := range rs {
			server.addResource(resources_path + r.Name())
		}

		// ranges through all resources
		for _, r := range server.resourceRoots {
			// ranges through all resource blocks
			for _, l := range r.Links() {
				blk, _ := server.blocks.GetBlock(ctx, l.Cid)
				log.Println(l.Cid, blk.RawData())
			}
		}

	}()

	select {}
}
