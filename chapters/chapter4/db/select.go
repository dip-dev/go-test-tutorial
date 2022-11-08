package db

import (
	"context"
	"fmt"

	"github.com/dip-dev/go-test-tutorial/mysql"
	"github.com/dip-dev/go-test-tutorial/mysql/queries"
	"github.com/dip-dev/go-test-tutorial/mysql/structs"
)

// DB ...
type DB struct {
	mysqlCli *mysql.Client
	queries  queries.Selecters
}

// New ...
func New(mysqlCli *mysql.Client, queries queries.Selecters) *DB {
	m := &DB{
		mysqlCli: mysqlCli,
		queries:  queries}
	return m
}

// SelectPrefectures ..
func (db *DB) SelectPrefectures(ctx context.Context, args map[string]interface{}) ([]structs.MPrefecture, error) {

	query := db.queries.SelectPrefectures()
	namedStmt, err := db.mysqlCli.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("m_prefecture取得 %s", err)
	}
	defer namedStmt.Close()

	var rows []structs.MPrefecture
	if err := namedStmt.SelectContext(ctx, &rows, args); err != nil {
		return nil, fmt.Errorf("m_prefecture取得 %s", err)
	}

	return rows, nil
}
