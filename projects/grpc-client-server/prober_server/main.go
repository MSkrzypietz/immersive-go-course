package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/CodeYourFuture/immersive-go-course/grpc-client-server/prober"
	"google.golang.org/grpc"
)

var (
	port            = flag.Int("port", 50051, "The server port")
	avgLatencyGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "avg_latency_gauge"}, []string{"endpoint"})
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
	avgLatencyMsecs := totalElapsedMsecs / float32(in.GetCount())
	avgLatencyGauge.WithLabelValues(in.GetEndpoint()).Set(float64(avgLatencyMsecs))

	return &pb.ProbeReply{
		AvgLatencyMsecs: avgLatencyMsecs,
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

	grpcServer := grpc.NewServer()
	pb.RegisterProberServer(grpcServer, &server{})
	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		fmt.Println("Starting Prometheus metrics server on :8080/metrics...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Failed to serve metrics: %v\n", err)
		}
	}()

	// Handle OS signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Stop the gRPC server gracefully
	fmt.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()

	fmt.Println("Server stopped")
}
