package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
	"github.com/jenniekibiri/Golang-psql-k8s-Banking-app/util"

)



var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err :=util.LoadConfig("../..")
	if err !=nil{
		log.Fatal("cannot load config: ", err)
	}
	
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	// use connecttion to c reate new testQueries object
	testQueries = New(testDB)
	// run the tests and exit
	os.Exit(m.Run())
}
