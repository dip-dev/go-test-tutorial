package chapter1

import "fmt"

func addtion(numA int, numB int) (int, error) {
	if numA < 0 {
		return 0, fmt.Errorf("numAは0以上の数値を指定してください。")
	}
	if numA > 100 {
		return 0, fmt.Errorf("numAは100以下の数値を指定してください。")
	}
	if numB < 10 {
		return 0, fmt.Errorf("numBは10以上の数値を指定してください。")
	}
	if numB > 200 {
		return 0, fmt.Errorf("numBは200以下の数値を指定してください。")
	}

	return numA + numB, nil
}
