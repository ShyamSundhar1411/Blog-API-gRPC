package main

import (
	"context"
	"log"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
)

func UpdateBlog(c pb.BlogServiceClient,id string){
	log.Println("Create Blog was invoked")
	newBlog := &pb.Blog{
		Id:id,
		AuthorId: "Not Test1",
		Title:"New Title",
		Content:"Content of first blog with update",
	}
	_,err:=c.UpdateBlog(context.Background(),newBlog)
	if err!=nil{
		log.Fatalf("Error while updating: %v\n",err)
	}
	log.Println("Blog was updated")
	
}