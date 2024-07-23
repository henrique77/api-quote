package main

import (
	"github.com/henrique77/api-quote/config/server"
	_ "github.com/henrique77/api-quote/docs"
)

// @title			API Quote
// @version		1.0
// @description	API responsible for managing freight quotes
// @termsOfService	http://swagger.io/terms/
func main() {
	server.New().Config().Start()
}
