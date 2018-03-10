// Package Dingo is the main entry point to your application.
package Dingo

import (
	"fmt"
	"github.com/dinever/golf"
	"github.com/dingoblog/dingo/app/handler"
	"github.com/dingoblog/dingo/app/model"
)



// Init loads a public and private key pair used to create and validate JSON
// web tokens, or creates a new pair if they don't exist. It also initializes
// the database connection.
func Init(privKey, pubKey string) {
	model.InitializeKey(privKey, pubKey)
	model.InitializeDb()
}

// Run starts our HTTP server on the given port.
func Run(portNumber string) {
	app := golf.New()
	app = handler.Initialize(app)
	fmt.Printf("Application Started on port %s\n", portNumber)
	app.Run(":" + portNumber)
}
