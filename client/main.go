package main

import (
	"context"
	"flag"
	"log"

	"github.com/heyts/minigrpc/pb"
	"google.golang.org/grpc"
)

var name = flag.String("name", "", "Name of person to greet")

func SayHi(client pb.GreetingClient, n string) {
	req := &pb.GreetingRequest{
		Person: n,
	}
	reply, _ := client.SayHi(context.Background(), req)
	log.Printf(reply.Message)
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	client := pb.NewGreetingClient(conn)

	SayHi(client, *name)

}
