package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/udev-21/iman-test-api-gateway/pkg/post-crud/pb"
)

func UpdatePost(ctx *gin.Context, c pb.PostCRUDServiceClient) {
	body := pb.Post{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if id != body.ID {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("request body id and url path id doesn't match"))
		return
	}

	// userId, _ := ctx.Get("userId")

	input := &pb.UpdatePostRequest{}
	input.Post = new(pb.Post)
	input.Post = &body
	res, err := c.Update(context.Background(), input)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusNoContent, &res)
}
