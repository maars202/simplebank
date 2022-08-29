package db

import (
	"database/sql"
	"log"
	"testing"
	"os"

	_ "github.com/lib/pq"
	// _ "github.com/k8-proxy/k8-go-comm"
)
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)
var testQueries *Queries
// create new global variables for db txs testing from store.go file:
var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	// conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err);
	}

	testQueries = New(testDB);
	// testQueries = New(conn);

	os.Exit(m.Run());
}