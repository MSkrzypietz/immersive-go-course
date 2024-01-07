package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/CodeYourFuture/immersive-go-course/grpc-client-server/prober"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement prober.ProberServer.
type server struct {
	pb.UnimplementedProberServer
}

func (s *server) DoProbes(ctx context.Context, in *pb.ProbeRequest) (*pb.ProbeReply, error) {
	var totalElapsed time.Duration
	var errorCount int32
	for i := 0; i < int(in.GetCount()); i++ {
		start := time.Now()
		resp, err := http.Get(in.GetEndpoint())
		if err != nil || resp.StatusCode != http.StatusOK {
			errorCount++
		} else {
			totalElapsed += time.Since(start)
		}
	}

	totalElapsedMsecs := float32(totalElapsed / time.Millisecond)
	return &pb.ProbeReply{
		AvgLatencyMsecs: totalElapsedMsecs / float32(in.GetCount()),
		SuccessCount:    in.GetCount() - errorCount,
		ErrorCount:      errorCount,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProberServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
