package chapter2

import (
	"fmt"
	"testing"
)

func TestExecWithMock(t *testing.T) {
	success := map[string]struct {
		want string
	}{
		"mock化テスト": {
			want: "Nice to meet you!!",
		},
	}

	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			fmt.Println(tc.want) // このfmtは一時的なエラー対応のため、テスト実装後に削除してください。

			// FIXME: mock contoroller生成

			// FIXME: mock設定

			// FIXME: mockを構造体(Chapter2)に設定

			// FIXME: exec関数を呼び出し

			// FIXME: 結果を検証

		})
	}
}
