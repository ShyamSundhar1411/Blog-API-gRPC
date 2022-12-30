package main

import (
	"log"
	"net"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var address string = "0.0.0.:50051"

type Server struct{
	pb.BlogServiceServer
}
func main(){
	listener,err:=net.Listen("tcp",address)
	if err!=nil{
		log.Fatalf("Failed to listen %v\n",err)
	}
	log.Printf("Listening on %s\n",address)
	serverInstance := grpc.NewServer()
	pb.RegisterBlogServiceServer(serverInstance,&Server{})
	reflection.Register(serverInstance)

	if err=serverInstance.Serve(listener);err!=nil{
		log.Fatalf("Failed to server: %v\n",err)
	}
}