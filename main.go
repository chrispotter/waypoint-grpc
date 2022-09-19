package main

import (
	"context"
	"crypto/tls"
	"log"

	pb "github.com/chrispotter/waypoint-grpc/pkg/server/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	metadata "google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	connConfig := tls.Config{
		InsecureSkipVerify: true,
	}
	creds := credentials.NewTLS(&connConfig)
	conn, err := grpc.DialContext(context.Background(), "localhost:9701", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "XNjgksMqa4Vf7BKct7M3jMv6BDRPyCPYgMLFomtXyw7eFFexFN8gjtdaS8vee2F7MfxEYHeFUBgjbWfHDb99YVzobQ6sHNnRt5TE7pP3PiuYcWpnHnUZbBFEPDBFpi23S2kXarga7anwgk")

	ctx = metadata.AppendToOutgoingContext(ctx, "client-api-protocol", "1,1")

	client := pb.NewWaypointClient(conn)
	versionInfo, err := client.ListProjects(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("failed to get version information: %v", err)
	}
	log.Printf("version information: %v", versionInfo)
}
