package main

import (
	"log"
	"net"
	"context"
	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var collection *mongo.Collection
var address string = "0.0.0.:50051"

type Server struct{
	pb.BlogServiceServer
}
func main(){
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err!=nil{
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err!=nil{
		log.Fatal(err)
	}
	collection = client.Database("blogdb").Collection("blog")
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