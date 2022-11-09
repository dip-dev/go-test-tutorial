package db

import (
	"log"
	"os"
	"testing"

	"github.com/dip-dev/go-test-tutorial/mysql"
)

var mysqlCli *mysql.Client

func TestMain(m *testing.M) {
	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %+v", err)
	}
	mysqlCli = cli

	os.Exit(m.Run())
}

// FIXME: ↓↓テストを作成する↓↓
