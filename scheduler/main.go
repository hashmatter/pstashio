package main

import (
	pb "github.com/hashmatter/pstashio/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	serverCtx := newServerCtx()
	pb.RegisterSchedulerServer(s, serverCtx)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal("failed to serve %v", err)
		}
	}()

	log.Println("Server listening on port", port)
	select {}
}
