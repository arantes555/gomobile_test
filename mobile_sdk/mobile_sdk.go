package gomobile_test

import (
	"fmt"
)

func Add(a int, b int) (int, error) {
	if a == 42 && b == 42 {
		return 0, fmt.Errorf("cannot add 42 with itself, it breaks the universe")
	}
	return a + b, nil
}
