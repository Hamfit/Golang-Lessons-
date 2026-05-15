package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
	t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// below is how a benchmark is structured
func BenchmarkRepeat(b *testing.B) {
	//... setup ...
	for b.Loop() {
		Repeat("a", 5) // code being measured
	}

	//... cleanup ...
}

func ExampleRepeat() {
	repeated := Repeat("Fawaz", 5)
	fmt.Println(repeated)
}