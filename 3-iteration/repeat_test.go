package repeat

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("with 5 repeat", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
func ExampleRepeat() {
	fmt.Println(Repeat("r", 5))
	// Output: rrrrr
}
