package main

import (
	"context"
	"log"

	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
)

func createBlog(c pb.BlogServiceClient)string{
	log.Println("Create Blog was invoked")
	blog := &pb.Blog{
		AuthorId: "Test_1",
		Title: "My First Blog",
		Content: "Content of First Blog",
	}
	response,err:=c.CreateBlog(context.Background(),blog)
	if err!=nil{
		log.Fatalf("Error %v\n",err)
	}
	log.Printf("Blog has been created %s\n",response.Id)
	return response.Id
}
