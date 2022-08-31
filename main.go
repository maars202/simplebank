// package-imports/main-package/main.go
package main

import (
	// "main"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/maars202/simplebank/api"
	db "github.com/maars202/simplebank/db/sqlc"
)

func init() {
    fmt.Println("launch initialization")
}

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
    // fmt.Println("launch the program !")
	conn, err := sql.Open(dbDriver, dbSource)
	// conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err);
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil{
		log.Fatal("cannot start server:", err)
	}


}