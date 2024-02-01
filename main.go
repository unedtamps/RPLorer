package main

import (
	"github.com/unedtamps/go-backend/bootstrap"
	"github.com/unedtamps/go-backend/util"
)

func main() {
	server, err := bootstrap.InitServer()
	if err != nil {
		util.Log.Fatal(err)
	}
	if err := server.Start(); err != nil {
		util.Log.Fatal(err)
	}
}
