package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlogs(c pb.BlogServiceClient){
	log.Println("List Blog was invoked")
	stream,err:=c.ListBlogs(context.Background(),&emptypb.Empty{})
	if err!=nil{
		log.Fatalf("Error while calling ListBlog: %v\n",err)
	}
	for{
		res,err:=stream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("Something happened %v\n",err)
		}
		log.Println(res)
	}

}