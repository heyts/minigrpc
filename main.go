// grpc minimal server
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/heyts/minigrpc/pb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const grpcAddr = ":9090"

// GreetingsServer is a GreetingsServer. Duh.
type GreetingsServer struct{}

// SayHi takes a context and a GreetingRequest and returns a GreetingReply
func (g *GreetingsServer) SayHi(ctx context.Context, r *pb.GreetingRequest) (*pb.GreetingReply, error) {
	msg := fmt.Sprintf("Hello, %s", r.Person)
	log.Printf("Received a name to greet: %s", r.Person)
	return &pb.GreetingReply{
		Message: msg,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to Listen to %s: %v", grpcAddr, err)
	}

	Server := grpc.NewServer()
	pb.RegisterGreetingServer(Server, &GreetingsServer{})
	log.Printf("GRPC Server Listening on: %s", grpcAddr)
	Server.Serve(listener)
}
