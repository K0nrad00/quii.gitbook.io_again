package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("got %q, want %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

// For some reasong Examples are not working for me
func ExampleRepeat(c string) {
	empty_str := ""
	for i := 0; i < 7; i++ {
		empty_str += c

	}
	fmt.Println(empty_str)
}
