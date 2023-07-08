package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

var dbDriver = "postgres"
var dbSource = "postgresql://root:pablo2846@localhost:5432/gopg-db-test?sslmode=disable"

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
