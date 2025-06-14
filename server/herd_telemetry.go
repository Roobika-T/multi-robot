package server

import (
	"fmt"
	"io"

	pb "headtelemetry/proto/v1/fleet_manager"
)

// HerdTelemetryServer implements the HerdTelemetry service
type HerdTelemetryServer struct {
	pb.UnimplementedFleetManagerServer
}

// NewHerdTelemetryServer creates a new instance of HerdTelemetryServer
func NewHerdTelemetryServer() *HerdTelemetryServer {
	return &HerdTelemetryServer{}
}

// HerdTelemetry implements the HerdTelemetry RPC method
func (s *HerdTelemetryServer) HerdTelemetry(stream pb.FleetManager_HerdTelemetryServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error receiving request: %v", err)
		}

		// Process the telemetry request based on the action
		response := &pb.HerdTelemetryResponse{
			SlugName:     req.SlugName,
			Action:       req.Action,
			Data:         req.Data,
			Status:       true,
			Message:      "Successfully processed telemetry data",
			Organization: req.Organization,
		}

		// Send the response back to the client
		if err := stream.SendAndClose(response); err != nil {
			return fmt.Errorf("error sending response: %v", err)
		}
	}
}

// Example usage of handling different telemetry actions
func processTelemetryAction(req *pb.HerdTelemetryReqest) (*pb.HerdTelemetryResponse, error) {
	switch req.Action {
	case pb.HerdTelemetryActions_FLEET_STATES:
		return handleFleetStates(req)
	case pb.HerdTelemetryActions_DISAPTCH_STATES:
		return handleDispatchStates(req)
	default:
		return &pb.HerdTelemetryResponse{
			SlugName:     req.SlugName,
			Action:       req.Action,
			Status:       false,
			Message:      "Unknown telemetry action",
			Organization: req.Organization,
		}, nil
	}
}

func handleFleetStates(req *pb.HerdTelemetryReqest) (*pb.HerdTelemetryResponse, error) {
	// Implement fleet states handling logic here
	return &pb.HerdTelemetryResponse{
		SlugName:     req.SlugName,
		Action:       req.Action,
		Data:         req.Data,
		Status:       true,
		Message:      "Fleet states processed successfully",
		Organization: req.Organization,
	}, nil
}

func handleDispatchStates(req *pb.HerdTelemetryReqest) (*pb.HerdTelemetryResponse, error) {
	// Implement dispatch states handling logic here
	return &pb.HerdTelemetryResponse{
		SlugName:     req.SlugName,
		Action:       req.Action,
		Data:         req.Data,
		Status:       true,
		Message:      "Dispatch states processed successfully",
		Organization: req.Organization,
	}, nil
}
