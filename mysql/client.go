package mysql

import (
	"fmt"

	// mysqlのドライバをインポート
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	driver string = "mysql"
	src    string = "TUTORIAL_USER:GO_TEST_TUTORIAL@tcp(mysql:3306)/TUTORIAL_USER?charset=utf8&parseTime=true&loc=Local"
)

// Client ..
type Client struct {
	DB *sqlx.DB
}

// New ..
func New() (*Client, error) {
	db, err := sqlx.Open(driver, src)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	client := &Client{
		DB: db,
	}
	return client, nil
}

// Close ..
func (c *Client) Close() {
	c.DB.Close()
}
