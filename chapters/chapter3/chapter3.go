package chapter3

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dip-dev/go-test-tutorial/mysql"
	"github.com/dip-dev/go-test-tutorial/mysql/queries"
	"github.com/dip-dev/go-test-tutorial/mysql/structs"
)

// Chapter3 ...
type Chapter3 struct {
	mysqlCli *mysql.Client
	queries  queries.Selecters
}

// New ...
func New(mysqlCli *mysql.Client, queries queries.Selecters) *Chapter3 {
	m := &Chapter3{
		mysqlCli: mysqlCli,
		queries:  queries}
	return m
}

// selectPrefecture ..
func (c3 *Chapter3) selectPrefecture(name string) (*structs.MPrefecture, error) {
	ctx := context.Background()

	args := map[string]interface{}{
		"name": name,
	}

	query := c3.queries.SelectPrefecture()
	namedStmt, err := c3.mysqlCli.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("m_prefecture取得 %s", err)
	}
	defer namedStmt.Close()

	var row structs.MPrefecture

	if err = namedStmt.GetContext(ctx, &row, args); err == sql.ErrNoRows {
		// sql.ErrNoRows以外のエラーはハンドリングしない
		return nil, nil
	}

	return &row, nil
}
