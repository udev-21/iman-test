package routes

import (
	"context"
	"net/http"

	"github.com/udev-21/iman-test-api-gateway/pkg/post-crud/pb"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context, c pb.PostCRUDServiceClient) {
	body := pb.Post{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// userId, _ := ctx.Get("userId")

	res, err := c.Create(context.Background(), &pb.CreatePostRequest{
		Post: &body,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
