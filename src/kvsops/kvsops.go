package kvsops

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hamorrar/Distributed-Key-Value-Store/src/kvs"
)

func PutKey(c *gin.Context) {
	key := c.Param("key")
	environmentVariable := strings.TrimSpace(os.Getenv("FORWARDING_ADDRESS"))

	// Handles logic for requests that need to be forwarded
	if environmentVariable != "" {
		address := "http://" + environmentVariable + "/kvs/" + key
		fmt.Println(address)

		client := http.Client{Timeout: time.Second * 10}

		req, err := http.NewRequest(http.MethodPut, address, c.Request.Body)

		if err != nil {
			return
		}

		req.Header.Set("Content-Type", "application/json")

		// Send the created HTTP request
		resp, respErr := client.Do(req)
		if respErr != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Cannot forward request"})
			return
		}
		defer resp.Body.Close()

		// Read the response from the server that received the forwarded request
		respBodyFromForward, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return
		}

		// Create variable to unmarshal the response body into
		var jsonBodyFromForward map[string]string
		err2 := json.Unmarshal(respBodyFromForward, &jsonBodyFromForward)
		if err2 != nil {
			return
		}

		c.JSON(resp.StatusCode, jsonBodyFromForward)
	}

	keyLength := len(key)
	if keyLength > 50 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is too long"})
		return
	}

	var jsonRequestBody map[string]interface{}
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	if e := json.Unmarshal(data, &jsonRequestBody); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}
	_, isProperRequest := jsonRequestBody["value"]
	if !isProperRequest {
		c.JSON(400, gin.H{"error": "PUT request does not specify a value"})
		return
	}
	_, found := kvs.KVS[key]
	if found {
		kvs.KVS[key] = jsonRequestBody["value"]
		c.JSON(http.StatusOK, gin.H{"result": "replaced"})
	} else {
		kvs.KVS[key] = jsonRequestBody["value"]
		c.JSON(http.StatusCreated, gin.H{"result": "created"})
	}

	fmt.Println(kvs.KVS)

}

func NoKey(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "No key specified."})
}

func GetKey(c *gin.Context) {
	key := c.Param("key")
	environmentVariable := strings.TrimSpace(os.Getenv("FORWARDING_ADDRESS"))
	if environmentVariable != "" {
		address := "http://" + environmentVariable + "/kvs/" + key

		fmt.Println(address)
	}
}
