package main

import (
	"log"

	"github.com/lnvn/workapp/Gomicro/internal/database"
	"github.com/lnvn/workapp/Gomicro/internal/server"
)

func main() {
	db,err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Fail to initialize Database client: %s", err)
	}
	srv :=server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
