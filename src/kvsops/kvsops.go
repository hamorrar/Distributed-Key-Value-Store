package kvsops

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hamorrar/Distributed-Key-Value-Store/src/kvs"
)

// Handler for routes without a key specified in URL.
func NoKey(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "No key specified."})
}

func PutKey(c *gin.Context) {
	key := c.Param("key")
	environmentVariable := strings.TrimSpace(os.Getenv("FORWARDING_ADDRESS"))

	// Handles logic for requests that need to be forwarded
	if environmentVariable != "" {
		forwardRequest(c, environmentVariable, http.MethodPut, key)
		return
	}

	keyLength := len(key)
	if keyLength > 50 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is too long"})
		return
	}

	var jsonRequestBody map[string]interface{}
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "PutKey cannot read request body"})
		return
	}

	if e := json.Unmarshal(data, &jsonRequestBody); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}
	_, isProperRequest := jsonRequestBody["value"]
	if !isProperRequest {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PUT request does not specify a value"})
		return
	}
	_, found := kvs.KVS[key]
	kvs.KVS[key] = jsonRequestBody["value"]
	if found {
		c.JSON(http.StatusOK, gin.H{"result": "replaced"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"result": "created"})
	}

	fmt.Println("---PUT KEY VALUE STORE: ", kvs.KVS)
}

func GetKey(c *gin.Context) {
	key := c.Param("key")
	environmentVariable := strings.TrimSpace(os.Getenv("FORWARDING_ADDRESS"))
	if environmentVariable != "" {
		fmt.Println("GET FORWARD")
		forwardRequest(c, environmentVariable, http.MethodGet, key)
		return
	}

	val, found := kvs.KVS[key]
	if found {
		fmt.Println("GET FOUND. KEY: ", key, "VALUE: ", val)
		c.JSON(http.StatusOK, gin.H{"result": "found", "value": val})
	} else {
		fmt.Println("GET NOT FOUND")
		c.JSON(http.StatusNotFound, gin.H{"error": "Key does not exist"})
	}
	fmt.Println("---GET KEY VALUE STORE: ", kvs.KVS)
}

func DeleteKey(c *gin.Context) {
	key := c.Param("key")
	environmentVariable := strings.TrimSpace(os.Getenv("FORWARDING_ADDRESS"))
	if environmentVariable != "" {
		fmt.Println("DELETE FORWARD")
		forwardRequest(c, environmentVariable, http.MethodDelete, key)
		return
	}

	for k := range kvs.KVS {
		if k == key {
			delete(kvs.KVS, k)
			c.JSON(http.StatusOK, gin.H{"result": "deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Key does not exist"})
}
