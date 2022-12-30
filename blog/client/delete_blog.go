package main

import (
	"context"
	"log"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
)

func DeleteBlog(c pb.BlogServiceClient,id string){
	log.Println("Delete Blog was invoked")
	_,err:=c.DeleteBlog(context.Background(),&pb.BlogId{Id:id})
	if err!=nil{
		log.Fatalf("Error while deleting %v\n",err)
	}
	log.Printf("Deleted Blog Successfully")
}