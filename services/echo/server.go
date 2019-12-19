package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"

	pb "github.com/crdsdev/crds/services/echo/genproto"
)

const listenPort = "7000"

type echoService struct{}

func main() {
	port := listenPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	svc := new(echoService)

	srv := grpc.NewServer()
	log.Infof("starting echo service on tcp: %q", lis.Addr().String())
	pb.RegisterEchoServiceServer(srv, svc)
	err = srv.Serve(lis)
	log.Fatal(err)
}

func (es *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Output: req.GetInput(),
	}, nil
}
