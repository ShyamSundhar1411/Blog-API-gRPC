package main

import (
	"context"
	"log"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server)ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error){
	log.Printf("Reading was invoked with: %v\n",in)
	oid,err:=primitive.ObjectIDFromHex(in.Id)
	if err!=nil{
		return nil,status.Errorf(
			codes.InvalidArgument,
			"Cannot Parse ID",
		)
	}
	data := &BlogItem{}
	filter := bson.M{"_id":oid}
	response := collection.FindOne(ctx,filter)
	if err := response.Decode(data);err!=nil{
		return nil,status.Errorf(
			codes.NotFound,
			"Cannot find blog with Id provided",
		)
	}
	return documentToBlog(data),nil
}