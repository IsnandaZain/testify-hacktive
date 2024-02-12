package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	How to run :
	- go test ./helper (package)
	- go test -v ./helper (package)
*/

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	// if result != 40 {
	// 	t.FailNow()
	// }

	if result != 40 {
		t.Error("Result has to be 40")
	}

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		panic("Result should be 20")
	}
}

func TestFailSumTestify(t *testing.T) {
	result := Sum(10, 10)

	require.Equal(t, 40, result, "Result has to be 40")
	fmt.Println("TestSum Eksekusi Terhenti")
}

func TestSumTestify(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")
	fmt.Println("TestSum Eksekusi Terhenti")
}
