// Package main implements a client for Prober service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/CodeYourFuture/immersive-go-course/grpc-client-server/prober"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr          = flag.String("addr", "localhost:50051", "the address to connect to")
	endpoint      = flag.String("endpoint", "http://www.google.com", "the endpoint to probe against")
	requestsCount = flag.Int("n", 5, "the number of requests to perform")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProberClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DoProbes(ctx, &pb.ProbeRequest{Endpoint: *endpoint, Count: int32(*requestsCount)})
	if err != nil {
		log.Fatalf("could not probe: %v", err)
	}
	log.Printf("Average Response Time: %f\n", r.GetAvgLatencyMsecs())
	log.Printf("Success count: %d\n", r.GetSuccessCount())
	log.Printf("Error count: %d\n", r.GetErrorCount())
}
