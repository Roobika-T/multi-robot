package client

import (
	"context"
	"log"

	pb "headtelemetry/proto/v1/fleet_manager"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TelemetryClient struct {
	client pb.FleetManagerClient
	conn   *grpc.ClientConn
}

func NewTelemetryClient(address string) (*TelemetryClient, error) {
	// Set up connection to the server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewFleetManagerClient(conn)
	return &TelemetryClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *TelemetryClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *TelemetryClient) SendTelemetry(ctx context.Context, slugName, data, organization string, action pb.HerdTelemetryActions) error {
	// Create a stream
	stream, err := c.client.HerdTelemetry(ctx)
	if err != nil {
		return err
	}

	// Create and send the request
	req := &pb.HerdTelemetryReqest{
		SlugName:     slugName,
		Action:       action,
		Data:         data,
		Organization: organization,
	}

	if err := stream.Send(req); err != nil {
		return err
	}

	// Close the stream and get the response
	response, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("Received response: %v", response)
	return nil
}
