package main

import (
	"Homework_56/genproto/pb"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c:=pb.NewStringServiceClient(conn)

	ctx, cancel :=context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Call tolower
	r, err := c.ToUpper(ctx, &pb.StringRequest{Input: input})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ToUpper: %s", r.Output)

}
