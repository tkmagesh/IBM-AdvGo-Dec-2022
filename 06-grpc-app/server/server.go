package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-app/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type appService struct {
	proto.UnimplementedAppServiceServer
}

func (s *appService) Add(ctx context.Context, req *proto.AddRequest) (res *proto.AddResponse, err error) {
	x := req.GetX()
	y := req.GetY()
	if dl, ok := ctx.Deadline(); ok {
		fmt.Printf("Request received with deadline : %v\n", dl)
	}
	fmt.Printf("Processing [Add] x = %d and y = %d\n", x, y)
	time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Cancel signal received...")
		err = errors.New("operation cancelled")
		return
	default:
		result := x + y
		fmt.Printf("Responding [Add] x = %d y = %d and result = %d\n", x, y, result)
		res = &proto.AddResponse{
			Result: result,
		}
		return
	}
	return
}

func main() {
	asi := &appService{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
