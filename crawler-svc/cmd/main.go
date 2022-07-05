package main

import (
	"fmt"
	"log"
	"net"

	"github.com/udev-21/iman-test-crawler-svc/pkg/config"
	"github.com/udev-21/iman-test-crawler-svc/pkg/crawler"
	"github.com/udev-21/iman-test-crawler-svc/pkg/crawler/pb"
	postcrud "github.com/udev-21/iman-test-crawler-svc/pkg/post-crud"
	"google.golang.org/grpc"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Crawler Svc on", c.Port)

	s := crawler.Server{
		PostCRUD: postcrud.InitServiceClient(&c),
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCrawlerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
