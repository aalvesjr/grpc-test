package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/aalvesjr/grpc-test/pb-post"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("could not to listen to port %d: %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterPostsServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}

var posts = map[int32]*pb.Post{
	1: &pb.Post{
		Id:          1,
		Title:       "the first",
		Description: "a little description of first post",
		CreatedAt:   fmt.Sprintf("%s", time.Now()),
	},
	2: &pb.Post{
		Id:          2,
		Title:       "the second",
		Description: "a little description of second post",
		CreatedAt:   fmt.Sprintf("%s", time.Now()),
	},
}

type server struct{}

func (server) Get(c context.Context, params *pb.SearchParams) (*pb.Post, error) {
	log.Printf("PostID received: (%d)", params.Id)

	post, ok := posts[params.Id]
	if !ok {
		return nil, fmt.Errorf("Post ID: %d, not found", params.Id)
	}
	return post, nil
}
