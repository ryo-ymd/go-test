package main

import (
	"github.com/ryo-ymd/go-test/db"
	"github.com/ryo-ymd/go-test/server"
)

func main() {
	db.Init()
	server.Init()
}
