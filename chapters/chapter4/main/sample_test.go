package main

// 動作確認したい場合はコメントアウトを外してください。

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"

// 	"github.com/dip-dev/go-test-tutorial/chapters/chapter4/da"
// 	"github.com/dip-dev/go-test-tutorial/mysql/structs"
// )
//
// func TestSample(t *testing.T) {
// 	ctx := context.Background()
// 	success := map[string]struct {
// 		area *string
// 		want Response
// 	}{
// 		"正常系サンプル": {
// 			area: strToPtr("北海道"),
// 			want: Response{
// 				Count: 1,
// 				Prefectures: []structs.MPrefecture{{
// 					ID:       1,
// 					Name:     "北海道",
// 					Area:     "北海道",
// 					LandArea: 83424,
// 				}},
// 			},
// 		},
// 	}
// 	fail := map[string]struct {
// 		area *string
// 		want string
// 	}{
// 		"異常系サンプル": {
// 			area: nil,
// 			want: "areaは必須です。",
// 		},
// 	}

// 	for tt, tc := range success {
// 		t.Run(tt, func(t *testing.T) {
// 			got, err := min.exec(ctx, tc.area)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tc.want, got)
// 		})
// 	}
// 	for tt, tc := range fail {
// 		t.Run(tt, func(t *testing.T) {
// 			got, err := min.exec(ctx, tc.area)
// 			assert.Nil(t, got)
// 			assert.Regexp(t, tc.want, err)
// 		})
// 	}
// }

// func TestMockSample(t *testing.T) {
// 	ctx := context.Background()
// 	success := map[string]struct {
// 		area    *string
// 		mockRes []structs.MPrefecture
// 		want    Response
// 	}{
// 		"正常系（モック）サンプル": {
// 			area: strToPtr("北海道"),
// 			mockRes: []structs.MPrefecture{{
// 				ID:       1,
// 				Name:     "北海道",
// 				Area:     "北海道",
// 				LandArea: 83424,
// 			}},
// 			want: Response{
// 				Count: 1,
// 				Prefectures: []structs.MPrefecture{{
// 					ID:       1,
// 					Name:     "北海道",
// 					Area:     "北海道",
// 					LandArea: 83424,
// 				}},
// 			},
// 		},
// 	}
// 	fail := map[string]struct {
// 		area    *string
// 		mockErr error
// 		want    string
// 	}{
// 		"異常系（モック）サンプル": {
// 			area:    strToPtr("北海道"),
// 			mockErr: fmt.Errorf("m_prefectures取得でエラーが発生しました。"),
// 			want:    "m_prefectures取得でエラーが発生しました。",
// 		},
// 	}

// 	for tt, tc := range success {
// 		t.Run(tt, func(t *testing.T) {

// 			// mock contoroller生成
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			// mock設定。想定される引数と戻り値がある関数であれば戻り値をReturnメソッドで設定する
// 			// 設定時に引数へ指定したものと異なる値がテスト時に渡されるとエラーとなり、テストが失敗する。
// 			mock := da.NewMockSelecters(ctrl)
// 			args := map[string]interface{}{
// 				"area": *tc.area,
// 			}
// 			mock.EXPECT().SelectPrefectures(ctx, args).Return(tc.mockRes, nil)

// 			// mockしたappmodelをserviceに設定
// 			s := New(mock)

// 			got, err := s.exec(ctx, tc.area)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tc.want, got)
// 		})
// 	}
// 	for tt, tc := range fail {
// 		t.Run(tt, func(t *testing.T) {

// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mock := da.NewMockSelecters(ctrl)
// 			args := map[string]interface{}{
// 				"area": *tc.area,
// 			}
// 			mock.EXPECT().SelectPrefectures(ctx, args).Return(nil, tc.mockErr)

// 			s := New(mock)

// 			got, err := s.exec(ctx, tc.area)
// 			assert.Nil(t, got)
// 			assert.Regexp(t, tc.want, err)
// 		})
// 	}
// }
