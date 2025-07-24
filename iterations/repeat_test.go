package iterations

import "testing"

func TestAdder(t *testing.T) {

	t.Run("repeat function testing", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})


	t.Run("saying hello to people", func(t *testing.T) {
		got := ExampleRepeat("a",5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
}


func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
	for b.Loop() {
		ExampleRepeat("a",5)
	}
}

