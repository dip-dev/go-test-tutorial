package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/dip-dev/go-test-tutorial/mysql"
	"github.com/dip-dev/go-test-tutorial/mysql/queries"
	"github.com/dip-dev/go-test-tutorial/mysql/structs"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mysqlCli *mysql.Client

func TestMain(m *testing.M) {
	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %+v", err)
	}
	mysqlCli = cli

	os.Exit(m.Run())
}

// FIXME: ↓↓テストを作成する↓↓
func TestSample(t *testing.T) {
	ctx := context.Background()

	success := map[string]struct {
		area string
		want []structs.MPrefecture
	}{
		"正常": {
			area: "北海道",
			want: []structs.MPrefecture{{
				ID:       1,
				Name:     "北海道",
				Area:     "北海道",
				LandArea: 83424,
			}},
		},
	}

	fail := map[string]struct {
		area string
	}{
		"異常（mock）": {
			area: "北海道",
		},
	}

	da := New(mysqlCli, queries.New())
	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			got, _ := da.SelectPrefectures(ctx, tc.area)
			assert.Equal(t, tc.want, got)
		})
	}

	for tt, tc := range fail {
		t.Run(tt, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := queries.NewMockSelecters(ctrl)
			mock.EXPECT().SelectPrefectures().Return("select * from abcd")
			d := New(mysqlCli, mock)
			got, _ := d.SelectPrefectures(ctx, tc.area)
			assert.Nil(t, got)
		})
	}
}
