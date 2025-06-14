package main

import (
	"log"
	"net"

	pb "headtelemetry/proto/v1/fleet_manager"
	"headtelemetry/server"

	"google.golang.org/grpc"
)

func main() {
	// Create a TCP listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the HerdTelemetry service
	pb.RegisterFleetManagerServer(s, server.NewHerdTelemetryServer())

	log.Printf("Server listening on :50051")

	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
