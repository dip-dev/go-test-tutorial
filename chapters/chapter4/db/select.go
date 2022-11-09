package db

import (
	"context"

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
func (db *DB) SelectPrefectures(ctx context.Context, area string) ([]structs.MPrefecture, error) {

	args := map[string]interface{}{
		"area": area,
	}

	query := db.queries.SelectPrefectures()
	namedStmt, err := db.mysqlCli.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer namedStmt.Close()

	var rows []structs.MPrefecture
	err = namedStmt.SelectContext(ctx, &rows, args)

	return rows, err
}
