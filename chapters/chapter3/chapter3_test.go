package chapter3

import (
	"log"
	"os"
	"testing"

	"github.com/dip-dev/go-test-tutorial/mysql"
)

var mysqlCli *mysql.Client

// TestMain(m *testing.M)については README.md を参照してください。
func TestMain(m *testing.M) {
	// 前処理 start

	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %v\n", err)
	}
	mysqlCli = cli

	os.Exit(m.Run())
}

// FIXME: ↓↓テストを作成する↓↓
