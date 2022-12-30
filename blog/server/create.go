package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server)CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error){
	log.Printf("CreateBlog was invoked\n")
	data := BlogItem{
		AuthorId: in.AuthorId,
		Title: in.Title,
		Content:in.Content,
	}
	response,err:=collection.InsertOne(ctx,data)
	if err!=nil{
		return nil,status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error :%v\n",err),
		)
	}
	oid,ok := response.InsertedID.(primitive.ObjectID)
	if !ok{
		return nil,status.Errorf(
			codes.Internal,
			"Cannot Convert to OID",
		)
	}
	return &pb.BlogId{
		Id:oid.Hex(),
	},nil
}