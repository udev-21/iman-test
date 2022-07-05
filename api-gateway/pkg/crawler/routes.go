package crawler

import (
	"github.com/gin-gonic/gin"
	"github.com/udev-21/iman-test-api-gateway/pkg/auth"
	"github.com/udev-21/iman-test-api-gateway/pkg/config"
	"github.com/udev-21/iman-test-api-gateway/pkg/crawler/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/crawler")
	routes.Use(a.AuthRequired)
	routes.POST("/start", svc.CrawlerStart)
	routes.GET("/status", svc.CrawlerStatus)
}

func (svc *ServiceClient) CrawlerStatus(ctx *gin.Context) {
	routes.CrawlerStatus(ctx, svc.Client)
}

func (svc *ServiceClient) CrawlerStart(ctx *gin.Context) {
	routes.CrawlerStart(ctx, svc.Client)
}
