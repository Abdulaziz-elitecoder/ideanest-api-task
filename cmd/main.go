// cmd/main.go
package main

import (
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/utils"
)

func main() {
	// Set up the Gin router
	router := utils.SetupRouter()

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
