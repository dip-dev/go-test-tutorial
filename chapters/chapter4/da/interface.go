package da

import (
	"context"

	"github.com/dip-dev/go-test-tutorial/mysql/structs"
)

// Selecters ..
type Selecters interface {
	SelectPrefectures(ctx context.Context, args map[string]interface{}) ([]structs.MPrefecture, error)
}
