package data

import (
	"context"
	"fmt"

	"github.com/dip-dev/go-test-tutorial/mysql"
	"github.com/dip-dev/go-test-tutorial/mysql/queries"
	"github.com/dip-dev/go-test-tutorial/mysql/structs"
)

// Data ...
type Data struct {
	mysqlCli *mysql.Client
	queries  queries.Selecters
}

// New ...
func New(mysqlCli *mysql.Client, queries queries.Selecters) *Data {
	m := &Data{
		mysqlCli: mysqlCli,
		queries:  queries}
	return m
}

// SelectPrefectures ..
func (d *Data) SelectPrefectures(ctx context.Context, area string) ([]structs.MPrefecture, error) {

	args := map[string]interface{}{
		"area": area,
	}

	query := d.queries.SelectPrefectures()
	namedStmt, err := d.mysqlCli.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("m_prefecture取得 %s", err)
	}
	defer namedStmt.Close()

	var rows []structs.MPrefecture
	namedStmt.SelectContext(ctx, &rows, args)

	return rows, nil
}
