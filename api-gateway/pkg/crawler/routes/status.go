package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udev-21/iman-test-api-gateway/pkg/crawler/pb"
)

func CrawlerStatus(ctx *gin.Context, c pb.CrawlerServiceClient) {

	res, err := c.Status(context.Background(), &pb.StatusRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
