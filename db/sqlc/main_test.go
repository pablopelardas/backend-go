package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/pablopelardas/backend-go/utils"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// Load config
	config, err := utils.LoadTestConfig("../../")
	if err != nil {
		log.Fatal("cannot load test config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
