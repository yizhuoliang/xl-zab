package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "../sim"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50053, "The server 3 port")
)

type server struct {
	pb.UnimplementedSimulationServer
}

// implementation of BroadcastTxn handler
func (s *server) Broadcast(ctx context.Context, in *pb.BroadcastTxn) (*pb.AcknowledgementTxn, error) {
	log.Printf("Follower 3 received: %v", in.GetContent())
	return &pb.AcknowledgementTxn{Content: "I Acknowledge"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSimulationServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
