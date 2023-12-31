package iteration

import "testing"

func TestRepeat(t *testing.T){
	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated{
		t.Errorf("Expected '%q', got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 1; i< b.N; i++{
		Repeat("a")
	}
}