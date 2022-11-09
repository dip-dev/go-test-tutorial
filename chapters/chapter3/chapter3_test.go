package chapter3

import (
	"log"
	"os"
	"testing"

	"github.com/dip-dev/go-test-tutorial/mysql"
	"github.com/dip-dev/go-test-tutorial/mysql/queries"
	"github.com/dip-dev/go-test-tutorial/mysql/structs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mysqlCli *mysql.Client

// TestMain(m *testing.M)については README.md を参照してください
func TestMain(m *testing.M) {
	// 前処理 start

	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %+v", err)
	}
	mysqlCli = cli

	os.Exit(m.Run())
}

// FIXME: ↓↓テストを作成する↓↓
func TestSample(t *testing.T) {

	success := map[string]struct {
		name string
		want *structs.MPrefecture
	}{
		"正常": {
			name: "北海道",
			want: &structs.MPrefecture{
				ID:       1,
				Name:     "北海道",
				Area:     "北海道",
				LandArea: 83424,
			},
		},
		"0件": {
			name: "アメリカ",
			want: nil,
		},
	}

	fail := map[string]struct {
		name string
	}{
		"異常（mock）": {
			name: "北海道",
		},
	}

	da := New(mysqlCli, queries.New())
	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			got, _ := da.selectPrefecture(tc.name)
			assert.Equal(t, tc.want, got)
		})
	}

	for tt, tc := range fail {
		t.Run(tt, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := queries.NewMockSelecters(ctrl)
			mock.EXPECT().SelectPrefecture().Return("select * from abcd")
			d := New(mysqlCli, mock)
			got, _ := d.selectPrefecture(tc.name)
			assert.Nil(t, got)
		})
	}
}
