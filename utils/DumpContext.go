package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func DumpContext(c *gin.Context) {
	// Create a map to hold the context data
	contextData := map[string]interface{}{
		"Request": map[string]interface{}{
			"Method": c.Request.Method,
			"URL":    c.Request.URL.String(),
			"Header": c.Request.Header,
			"Body":   c.Request.Body,
		},
		"Writer": map[string]interface{}{
			"Status": c.Writer.Status(),
			"Header": c.Writer.Header(),
		},
		"Params": c.Params,
		"Keys":   c.Keys,
	}

	// Convert the map to JSON for easy logging
	contextJSON, err := json.MarshalIndent(contextData, "", "  ")
	if err != nil {
		log.Println("Error dumping context:", err)
		return
	}

	// Open the log file in append mode
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Create a new logger that writes to the file
	logger := log.New(file, "", log.LstdFlags)

	// Log the context data
	logger.Println("Context Dump:", string(contextJSON))
}
