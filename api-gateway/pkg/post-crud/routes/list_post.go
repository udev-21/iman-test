package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udev-21/iman-test-api-gateway/pkg/post-crud/pb"
)

func ListPost(ctx *gin.Context, c pb.PostCRUDServiceClient) {
	body := pb.ListPostRequest{}

	//no error checking because all fields optional
	ctx.BindJSON(&body)

	res, err := c.List(context.Background(), &body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
