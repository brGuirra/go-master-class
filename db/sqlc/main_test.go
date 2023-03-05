package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/brGuirra/simple-bank/utils"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
