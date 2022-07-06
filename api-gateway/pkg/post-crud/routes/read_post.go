package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/udev-21/iman-test-api-gateway/pkg/post-crud/pb"
)

func ReadPost(ctx *gin.Context, c pb.PostCRUDServiceClient) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	// userId, _ := ctx.Get("userId")

	res, err := c.Read(context.Background(), &pb.ReadPostRequest{
		ID: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
