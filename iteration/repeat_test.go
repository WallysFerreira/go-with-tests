package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 7)
	expected := "aaaaaaa"

	if got != expected {
		t.Errorf("Got %q expected %q", got, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
