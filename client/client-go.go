package main

import (
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/aalvesjr/grpc-test/pb-post"
	"golang.org/x/net/context"
)

func main() {
	fmt.Println("ok")
	backend := flag.String("b", "localhost:8080", "address of the Posts backend")
	id := flag.Uint("id", 0, "ID to search Posts in backend")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not to connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewPostsClient(conn)

	p, err := client.Get(context.Background(), &pb.SearchParams{Id: int32(*id)})
	if err != nil {
		log.Fatalf("error to find Post: %v", err)
	} else {
		fmt.Println(p)
	}
}
