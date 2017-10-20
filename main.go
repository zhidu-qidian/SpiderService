package main

import (
	"github.com/gin-gonic/gin"

	"workspace/SpiderService/config"
	"workspace/SpiderService/handlers"
	_ "workspace/SpiderService/storage"
)

func init() {
	if !config.C.Global.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	r := gin.Default()
	// store service group
	storeRouter := r.Group("/api/store")
	{
		storeRouter.POST("/video", handlers.StoreVideoHandler)
		storeRouter.POST("/joke", handlers.StoreJokeHandler)
		storeRouter.POST("/comment", handlers.StoreCommentHandler)
		storeRouter.POST("/news", handlers.StoreNewsHandler)
	}
	// update service group
	updateRouter := r.Group("/api/update")
	{
		updateRouter.PUT("/comment", handlers.UpdateCommentHandler)
	}
	r.Run(config.C.Global.Listen)
}
