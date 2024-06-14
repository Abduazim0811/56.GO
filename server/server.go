package main

import (
	"Homework_56/genproto/pb"
	"context"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedStringServiceServer
}

func (s *Server) ToUpper(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	return &pb.StringResponse{Output: strings.ToUpper(req.Input)}, nil
}

func (s *Server) ToLower(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	return &pb.StringResponse{Output: strings.ToLower(req.Input)}, nil
}

func (s *Server) Reverse(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	runes := []rune(req.Input)
	for i,j :=0, len(runes)-1; i<j; i,j=i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return &pb.StringResponse{Output: string(runes)}, nil
}

func (s *Server) Length(ctx context.Context, req *pb.StringRequest) (*pb.LengthResponse, error) {
	return &pb.LengthResponse{Length: int32(len(req.Input))}, nil
}


func main(){
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	
	s:=grpc.NewServer()
	pb.RegisterStringServiceServer(s,&Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err!=nil{
		log.Fatal(err)
	} 
}