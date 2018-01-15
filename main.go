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
		v1.POST("/:filename", addFile)
	}
	router.Run()
}
