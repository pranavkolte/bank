package db

import (
	"bank/util"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Failed to load conig", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	fmt.Println("Connecting to DB")
	if err != nil {
		log.Fatal("cannot connect to Database", err)
	}

	testQuries = New(testDB)

	// Run the tests
	os.Exit(m.Run())
}

func TestDBConnection(t *testing.T) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Failed to load conig", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		t.Fatalf("failed to open database connection: %v", err)
	}

	err = conn.Ping()
	if err != nil {
		t.Fatalf("failed to ping database: %v", err)
	}

	// Close the connection
	conn.Close()
}
