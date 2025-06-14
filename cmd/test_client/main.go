package main

import (
	"context"
	"log"
	"time"

	"headtelemetry/client"
	pb "headtelemetry/proto/v1/fleet_manager"
)

func main() {
	// Create a new telemetry client
	telemetryClient, err := client.NewTelemetryClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer telemetryClient.Close()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Example: Send fleet states telemetry
	err = telemetryClient.SendTelemetry(
		ctx,
		"test-slug",
		`{"fleet_id": "fleet-1", "status": "active", "robots": ["robot-1", "robot-2"]}`,
		"test-org",
		pb.HerdTelemetryActions_FLEET_STATES,
	)
	if err != nil {
		log.Fatalf("Failed to send fleet states telemetry: %v", err)
	}

	// Example: Send dispatch states telemetry
	err = telemetryClient.SendTelemetry(
		ctx,
		"test-slug",
		`{"dispatch_id": "dispatch-1", "status": "in_progress", "tasks": ["task-1", "task-2"]}`,
		"test-org",
		pb.HerdTelemetryActions_DISAPTCH_STATES,
	)
	if err != nil {
		log.Fatalf("Failed to send dispatch states telemetry: %v", err)
	}
}
