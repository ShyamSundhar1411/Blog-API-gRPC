package main

import (
	"context"
	"log"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
)

func ReadBlog(c pb.BlogServiceClient,id string) *pb.Blog{
	log.Println("Create Blog was invoked")
	
	request := &pb.BlogId{Id:id}
	response,err:=c.ReadBlog(context.Background(),request)
	if err!=nil{
		log.Printf("Error happend while reading:%v\n",err)
	}
	log.Printf("Blog was read:%v\n",response)
	return response
}