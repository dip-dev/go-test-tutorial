package chapter6

import (
	"fmt"
	"time"
)

func dispArg(in any) {
	time.Sleep(1 * time.Second)
	fmt.Println("Received:", in)
}
