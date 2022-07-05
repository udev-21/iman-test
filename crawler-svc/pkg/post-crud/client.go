package postcrud

import (
	"fmt"

	"github.com/udev-21/iman-test-crawler-svc/pkg/post-crud/pb"

	"github.com/udev-21/iman-test-crawler-svc/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.PostCRUDServiceClient
}

func InitServiceClient(c *config.Config) pb.PostCRUDServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.PostCRUDSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewPostCRUDServiceClient(cc)
}
