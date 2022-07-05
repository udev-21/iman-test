package main

import (
	"fmt"
	"log"
	"net"

	"github.com/udev-21/iman-test-post-crud-svc/pkg/config"
	"github.com/udev-21/iman-test-post-crud-svc/pkg/db"
	"github.com/udev-21/iman-test-post-crud-svc/pkg/pb"
	"github.com/udev-21/iman-test-post-crud-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPostCRUDServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
