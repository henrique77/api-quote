package main

import "github.com/henrique77/api-quote/config/server"

func main() {
	server.New().Config().Start()
}
