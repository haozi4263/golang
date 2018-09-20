package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"fmt"
)

func main()  {
	route := gin.Default()
	route.MaxMultipartMemory = 32 << 64
	route.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprintf("c:/tmp/%s", file.Filename)
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' upload!", file.Filename),
		})
	})
	route.Run()
}
