package utils

import (
	"fmt"
	"os"
	"testing"
)

func addition(a, b int, result *int) {
	*result = a + b
}

func TestPointer(t *testing.T) {
	var result int
	addition(3, 2, &result)

	fmt.Printf("Result: %d", result)
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
