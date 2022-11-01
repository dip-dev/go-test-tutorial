package main

import (
	"log"
	"os"
	"testing"

	data "github.com/dip-dev/go-test-tutorial/chapters/chapter4/da"
	"github.com/dip-dev/go-test-tutorial/mysql"
	"github.com/dip-dev/go-test-tutorial/mysql/queries"
)

var min *Main

func TestMain(m *testing.M) {
	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %v\n", err)
	}
	min = New(data.New(cli, queries.New()))
	os.Exit(m.Run())
}

// 文字列をポインタ型に変換する
func strToPtr(s string) *string {
	return &s
}

// FIXME: ↓↓テストを作成する↓↓
