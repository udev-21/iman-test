package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/udev-21/iman-test-api-gateway/pkg/auth"
	"github.com/udev-21/iman-test-api-gateway/pkg/config"
	"github.com/udev-21/iman-test-api-gateway/pkg/crawler"
	postcrud "github.com/udev-21/iman-test-api-gateway/pkg/post-crud"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	crawler.RegisterRoutes(r, &c, &authSvc)
	postcrud.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
