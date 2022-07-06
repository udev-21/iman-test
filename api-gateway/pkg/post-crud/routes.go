package postcrud

import (
	"github.com/gin-gonic/gin"
	"github.com/udev-21/iman-test-api-gateway/pkg/auth"
	"github.com/udev-21/iman-test-api-gateway/pkg/config"
	"github.com/udev-21/iman-test-api-gateway/pkg/post-crud/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/posts")
	routes.Use(a.AuthRequired)

	routes.POST("", svc.CreatePost)
	routes.GET("", svc.ListPost)
	routes.GET("/:id", svc.ReadPost)
	routes.PUT("/:id", svc.UpdatePost)

	routes.DELETE("/:id", svc.DeletePost)
}

func (svc *ServiceClient) CreatePost(ctx *gin.Context) {
	routes.CreatePost(ctx, svc.Client)
}

func (svc *ServiceClient) ReadPost(ctx *gin.Context) {
	routes.ReadPost(ctx, svc.Client)
}

func (svc *ServiceClient) DeletePost(ctx *gin.Context) {
	routes.DeletePost(ctx, svc.Client)
}

func (svc *ServiceClient) ListPost(ctx *gin.Context) {
	routes.ListPost(ctx, svc.Client)
}

func (svc *ServiceClient) UpdatePost(ctx *gin.Context) {
	routes.UpdatePost(ctx, svc.Client)
}
