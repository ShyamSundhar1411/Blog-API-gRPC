package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server)ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error{
	log.Printf("List Blog Function invoked\n")
	cur,err:=collection.Find(context.Background(),primitive.D{})
	if err!=nil{
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Internal error %v\n",err),
		)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()){
		data := &BlogItem{}
		err:=cur.Decode(data)
		if err!=nil{
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Unknown Internal error while decoding %v\n",err),
			)
		}
		stream.Send(documentToBlog(data))
	}
	if err=cur.Err();err!=nil{
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Internal error %v\n",err),
		)
	}
	return nil
}
