package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()
	// doRequestResponse(ctx, client)
	doServerStreaming(ctx, client)

}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	/*
		addCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		fmt.Println("Hit ENTER to stop...")
		go func() {
			fmt.Scanln()
			cancel()
		}()
	*/

	addCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := client.Add(addCtx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Result : %d\n", res.GetResult())
}

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All responses received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Prime Number received :", res.GetPrimeNo())
	}
}
