package da

import (
	"context"
	"fmt"

	"github.com/dip-dev/go-test-tutorial/mysql"
	"github.com/dip-dev/go-test-tutorial/mysql/queries"
	"github.com/dip-dev/go-test-tutorial/mysql/structs"
)

// DataAccess ...
type DataAccess struct {
	mysqlCli *mysql.Client
	queries  queries.Selecters
}

// New ...
func New(mysqlCli *mysql.Client, queries queries.Selecters) *DataAccess {
	m := &DataAccess{
		mysqlCli: mysqlCli,
		queries:  queries}
	return m
}

// SelectPrefectures ..
func (da *DataAccess) SelectPrefectures(ctx context.Context, args map[string]interface{}) ([]structs.MPrefecture, error) {

	query := da.queries.SelectPrefectures()
	namedStmt, err := da.mysqlCli.DB.PrepareNamedContext(ctx, query)
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
