package main

import (
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	server := newServerCtx()
	pb.RegisterSchedulerServer(s, server)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal("failed to serve", err)
		}
	}()

	log.Println("Server listening on port", port)

	// loads all the blocks into the server's blockstore asynchronously
	go func() {
		// get all files in resource path
		rs, _ := ioutil.ReadDir(resources_path)
		for _, r := range rs {
			server.addResource(resources_path + r.Name())
			log.Println("added", r.Name())
		}
	}()

	select {}
}
