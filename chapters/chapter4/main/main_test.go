package main

import (
	"log"
	"os"
	"testing"

	"github.com/dip-dev/go-test-tutorial/chapters/chapter4/da"
	"github.com/dip-dev/go-test-tutorial/mysql"
	"github.com/dip-dev/go-test-tutorial/mysql/queries"
)

var mainTest *Main

func TestMain(m *testing.M) {
	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %+v", err)
	}
	mainTest = New(da.New(cli, queries.New()))
	os.Exit(m.Run())
}

// 文字列をポインタ型に変換する
func strToPtr(t *testing.T, s string) *string {
	t.Helper()
	return &s
}

// FIXME: ↓↓テストを作成する↓↓
