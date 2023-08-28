package chapter7

import (
	"fmt"
	"os"
	"time"
	"unicode/utf8"
)

func createTextFile(name string) error {
	time.Sleep(1 * time.Second)
	if utf8.RuneCountInString(name) > 20 {
		return fmt.Errorf("invalid name length")
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fp, err := os.Create(fmt.Sprintf("%s/%s.txt", dir, name))
	if err != nil {
		return err
	}
	defer fp.Close()

	_, _ = fp.WriteString(fmt.Sprintf("create:%s.txt", name))

	return nil
}
