package main

import (
	"github.com/go-micro-v4-demo/helloworld/handler"
	pb "github.com/go-micro-v4-demo/helloworld/proto"
	userPb "github.com/go-micro-v4-demo/user/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "helloworld"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)
	client := srv.Client()
	helloworld := &handler.Helloworld{
		UserService: userPb.NewUserService("user", client),
	}
	if err := pb.RegisterHelloworldHandler(srv.Server(), helloworld); err != nil {
		logger.Fatal(err)
	}
	// Register handler
	//if err := pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld)); err != nil {
	//	logger.Fatal(err)
	//}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
