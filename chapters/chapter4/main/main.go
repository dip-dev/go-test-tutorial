package main

import (
	"context"
	"fmt"
	"unicode/utf8"

	"github.com/dip-dev/go-test-tutorial/chapters/chapter4/db"
	"github.com/dip-dev/go-test-tutorial/mysql/structs"
)

// Main ..
type Main struct {
	db db.Selecters
}

// Response ..
type Response struct {
	Count       int
	Prefectures []structs.MPrefecture
}

// New ..
func New(db db.Selecters) *Main {
	return &Main{
		db: db,
	}
}

func (m Main) exec(ctx context.Context, area *string) (interface{}, error) {

	if err := validate(area); err != nil {
		return nil, err
	}

	args := map[string]interface{}{
		"area": *area,
	}
	prefectures, err := m.db.SelectPrefectures(ctx, args)
	if err != nil {
		return nil, err
	}

	count := len(prefectures)
	if count == 0 {
		return Response{
			Count:       count,
			Prefectures: []structs.MPrefecture{},
		}, nil
	}

	return Response{
		Count:       count,
		Prefectures: prefectures,
	}, nil
}

func validate(area *string) error {
	if area == nil || *area == "" {
		return fmt.Errorf("areaは必須です。")
	}
	if utf8.RuneCountInString(*area) > 10 {
		return fmt.Errorf("areaは10文字以内で指定してください。")
	}

	return nil
}
