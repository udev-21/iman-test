package crawler

import (
	"fmt"

	"github.com/udev-21/iman-test-api-gateway/pkg/config"
	"github.com/udev-21/iman-test-api-gateway/pkg/crawler/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.CrawlerServiceClient
}

func InitServiceClient(c *config.Config) pb.CrawlerServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CrawlerSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCrawlerServiceClient(cc)
}
