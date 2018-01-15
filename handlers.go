package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func fetchFile(c *gin.Context) {
	filepath := FileAccessor.read(c.Param("filename"))
	c.File(filepath) 
}

func addFile(c *gin.Context) {
	result := FileAccessor.add(c.Param("filename"), c.Request.Body)
	if result == "error" {
		c.String(http.StatusInternalServerError, "error")
	} else {
		c.String(http.StatusOK, "ok")
	}
}

func updateFile(c *gin.Context) {
	result := FileAccessor.add(c.Param("filename"), c.Request.Body)
	if result == "error" {
		c.String(http.StatusInternalServerError, "error")
	} else {
		c.String(http.StatusOK, "ok")
	}
}

func deleteFile(c *gin.Context) {
	result := FileAccessor.delete(c.Param("filename"))
	if result == "error" {
		c.String(http.StatusInternalServerError, "error")
	} else {
		c.String(http.StatusOK, "ok")
	}
}
