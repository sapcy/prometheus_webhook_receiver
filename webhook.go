package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func receive(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	fmt.Println(c.Request.Header)
	fmt.Println(c.Request.Host)
	fmt.Println(string(body))
}
