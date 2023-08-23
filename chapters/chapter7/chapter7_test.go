package chapter7

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func deleteTextFile(t *testing.T, target string) {

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Remove(fmt.Sprintf("%s/%s.txt", dir, target)); err != nil {
		t.Fatal(err)
	}
}

func TestCreateTextFile(t *testing.T) {

	fileNames := []string{"file1", "file2", "file3"}

	t.Run("成功", func(t *testing.T) {
		for i, fileName := range fileNames {
			t.Run(fmt.Sprintf("create file%d", i), func(t *testing.T) {
				err := createTextFile(fileName)
				assert.NoError(t, err)
			})
			deleteTextFile(t, fileName)
		}
	})
}
