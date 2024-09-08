package main

import (
	"log"
	"net"

	pb "github.com/ibreakthecloud/traintkt/pkg/proto"
	"github.com/ibreakthecloud/traintkt/pkg/store/inmemory"
	"github.com/ibreakthecloud/traintkt/pkg/ticketservice"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialise a store: in-memory store
	// store is an interface, which later can be replaced with a database store
	store := inmemory.New()

	s := grpc.NewServer()
	pb.RegisterTicketServiceServer(s, ticketservice.NewTicketServer(store))

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
