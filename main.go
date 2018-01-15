package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"io"
	"path/filepath"
	"net/http"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/")
	{
		v1.GET("/:filename", fetchFile)
		v1.POST("/:filename", addFile)
	}
	router.Run()
}

func fetchFile(c *gin.Context) {
	filename := c.Param("filename")
	path     := config("dir")
	fullname := filepath.Join(path, filename)
	fmt.Println("%v", c.Param("filename"))
	c.File(fullname) 
}

func addFile(c *gin.Context) {
	filename := c.Param("filename")
	fmt.Println("%v", filename)
	path     := config("dir")
	fmt.Println("%v", path)
	fullname := filepath.Join(path, filename)
	fmt.Println("%v", fullname)
	content  := c.Request.Body
	fmt.Println("%v", content)
	file, err := os.Create(fullname)
	fmt.Println("%v", err)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	io.Copy(file, content)
	if r := recover(); r != nil {
		c.String(http.StatusOK, "error")
	}
	c.String(http.StatusOK, "ok")
	// return ""
}

func config(key string) string {
	switch key {
		case "dir":
			return "download/"
	}
	return ""
}