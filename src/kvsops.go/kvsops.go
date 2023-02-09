package kvsops

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PutKey(c *gin.Context) {
	fmt.Println("PUT KEY CALLED")
}
