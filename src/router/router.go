package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hamorrar/Distributed-Key-Value-Store/src/kvsops"
)

func InitRoutes(router *gin.Engine, KVS map[string]interface{}, environmentVariable string) {
	router.PUT("/kvs/:key", kvsops.PutKey)

	router.GET("/kvs/:key", kvsops.GetKey)

	router.Any("/kvs", kvsops.NoKey)
	router.Any("/kvs/", kvsops.NoKey)
}
