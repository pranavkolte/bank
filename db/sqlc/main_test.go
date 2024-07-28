package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/bank?sslmode=disable"
)

var testQuries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	fmt.Println("Connecting to DB")
	if err != nil {
		log.Fatal("cannot connect to Database", err)
	}

	testQuries = New(testDB)

	// Run the tests
	os.Exit(m.Run())
}

func TestDBConnection(t *testing.T) {
	conn, err := sql.Open(dbDriver, dbSource)
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
