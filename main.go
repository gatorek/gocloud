package main

import (
	"github.com/gin-gonic/gin"
)

var FileAccessor FileAcceess

func main() {
	FileAccessor = getFileAccessor(config("fastrategy"))
	router := gin.Default()
	v1 := router.Group("/")
	{
		v1.GET("/:filename", fetchFile)
		v1.PUT("/:filename", addFile)
		v1.POST("/:filename", updateFile)
		v1.DELETE("/:filename", deleteFile)
	}
	router.Run()
}
