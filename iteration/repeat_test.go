package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 7)
	expected := "aaaaaaa"

	if got != expected {
		t.Errorf("Got %q expected %q", got, expected)
	}
}

func TestReplace(t *testing.T) {
	got := strings.ReplaceAll("Testing one two three", "three", "tested")
	expected := "Testing one two tested"

	if got != expected {
		t.Errorf("Got %q expected %q", got, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("teste", 2)
	fmt.Println(repeated)
	// Output: testeteste
}
