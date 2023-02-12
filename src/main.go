package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hamorrar/Distributed-Key-Value-Store/src/kvs"
	"github.com/hamorrar/Distributed-Key-Value-Store/src/router"
)

func main() {

	KVS := kvs.InitKVS()

	environmentVariable := strings.TrimSpace(os.Getenv("FORWARDING_ADDRESS"))
	ginRouter := gin.Default()
	ginRouter.SetTrustedProxies(nil)

	router.InitRoutes(ginRouter, KVS, environmentVariable)

	ginRouter.Run(":8090")
}
