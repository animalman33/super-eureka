package routing

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var URL_PREFIX = os.Getenv("BASE_URL")

func StartServer() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	//router.TrustedPlatform = gin.PlatformCloudflare

	{
		apiGroup := router.Group("/api" + URL_PREFIX)
		apiGroup.GET("/health", HealthCheck)
	}

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
