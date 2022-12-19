package db

import (
	"context"

	"github.com/dip-dev/go-test-tutorial/mysql/structs"
)

// Selecters ..
type Selecters interface {
	SelectPrefectures(ctx context.Context, area string) ([]structs.MPrefecture, error)
}
