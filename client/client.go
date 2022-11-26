package main

import (
	"context"
	"google.golang.org/grpc"

	pb "github.com/mark4z/grpc-reflection/api"
	"google.golang.org/grpc/credentials/insecure"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pb.NewGreeterClient(conn)
	hello, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "mark4z",
		Pod: &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name: "podName",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	println(hello.Message)
}
