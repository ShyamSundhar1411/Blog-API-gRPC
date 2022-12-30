package main
import (
	pb "github.com/ShyamSundhar1411/Blog-gRPC/blog/proto"
)

type BlogItem struct{
	pb.BlogServiceServer
}