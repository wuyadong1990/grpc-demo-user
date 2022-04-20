package main

import (
	"context"
	"log"
	"time"

	pb "github.com/wuyadong1990/grpc-demo-proto/user"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:5000"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UserInfo(ctx, &pb.UserBase{
		UserName: "test",
		Iphone:   "18201420251",
		Sex:      1,
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetId())
}
