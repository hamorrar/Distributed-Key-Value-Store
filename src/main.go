package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hamorrar/Distributed-Key-Value-Store/src/router"
)

var kvs map[string]interface{}
var environmentVariable string = strings.TrimSpace(os.Getenv("FORWARDING_ADDRESS"))

func main() {
	kvs = make(map[string]interface{})

	ginRouter := gin.Default()
	ginRouter.SetTrustedProxies(nil)

	router.InitRoutes(ginRouter, kvs, environmentVariable)

	ginRouter.Run(":8090")
}
