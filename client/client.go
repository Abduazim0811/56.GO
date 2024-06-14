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
	addr := flag.String("addr", "localhost:12345", "the address to connect to")
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStringServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	input := "Hello, gRPC!"

	r, err := c.ToUpper(ctx, &pb.StringRequest{Input: input})
	if err != nil {
		log.Fatalf("could not convert to upper: %v", err)
	}
	log.Printf("ToUpper: %s", r.Output)

	r, err = c.ToLower(ctx, &pb.StringRequest{Input: input})
	if err != nil {
		log.Fatalf("could not convert to lower: %v", err)
	}
	log.Printf("ToLower: %s", r.Output)

	r, err = c.Reverse(ctx, &pb.StringRequest{Input: input})
	if err != nil {
		log.Fatalf("could not reverse: %v", err)
	}
	log.Printf("Reverse: %s", r.Output)

	lenResp, err := c.Length(ctx, &pb.StringRequest{Input: input})
	if err != nil {
		log.Fatalf("could not get length: %v", err)
	}
	log.Printf("Length: %d", lenResp.Length)
}
