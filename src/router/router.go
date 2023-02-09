package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hamorrar/Distributed-Key-Value-Store/src/kvsops.go"
)

func InitRoutes(router *gin.Engine, kvs map[string]interface{}, environmentVariable string) {
	router.PUT("/kvs/:key", kvsops.PutKey)
}
