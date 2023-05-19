package kvsops

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func forwardRequest(c *gin.Context, environmentVariable string, method string, key string) {
	address := "http://" + environmentVariable + "/kvs/" + key

	client := http.Client{Timeout: time.Second * 10}

	req, err := http.NewRequest(method, address, c.Request.Body)

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
	var jsonBodyFromForward map[string]interface{}
	err2 := json.Unmarshal(respBodyFromForward, &jsonBodyFromForward)
	if err2 != nil {
		return
	}

	fmt.Println("HELPER STATUS CODE: ", resp.StatusCode)
	fmt.Println("HELPER JSONBODYFROMFORWARD: ", jsonBodyFromForward)
	c.JSON(resp.StatusCode, jsonBodyFromForward)
}
