package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "../sim"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverPorts = [6]string{"", "localhost:50051", "localhost:50052",
	"localhost:50053", "localhost:50054", "localhost:50055"}

var (
	addr1 = flag.String("addr1", "localhost:50051", "the address of server1")
)

func main() {
	// build 4 client connections
	var clients [5]pb.SimulationClient
	for i := 1; i <= 4; i++ {
		conn, err := grpc.Dial(serverPorts[i], grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect server%d: %v", i, err)
		}
		defer conn.Close()
		clients[i] = pb.NewSimulationClient(conn)
	}

	// let's do a broadcast!
	akgCount := 0

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for i := 1; i <= 4; i++ {
		r, err := clients[i].Broadcast(ctx, &pb.BroadcastTxn{Content: "Leader is calling you!"})
		if err != nil {
			log.Fatalf("could not broadcast to server%d: %v", i, err)
		}
		if r.GetContent() == "I Acknowledge" {
			akgCount++
		}
	}
	log.Printf("Result: %d acknowledgement received", akgCount)
}

// func main() {

// 	flag.Parse()
// 	conn1, err := grpc.Dial(*addr1, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect server1: %v", err)
// 	}
// 	defer conn1.Close()
// 	c := pb.NewSimulationClient(conn1)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := c.Broadcast(ctx, &pb.BroadcastTxn{Content: "Leader is calling you!"})
// 	if err != nil {
// 		log.Fatalf("could not broadcast to server1", err)
// 	}
// 	log.Printf("Acknoledgement received: %s", r.GetContent())
// }
