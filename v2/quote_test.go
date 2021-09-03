package quote

import "testing"

func TestSay(t *testing.T) {
	want := "Hello"
	if got := Say2(); got != want {
		t.Errorf("Say() = %q , want : %q", got, want)
	}
}
func TestSpeak(t *testing.T) {
	want := "Hi,mate"
	if got := Sqeak2(); got != want {
		t.Errorf("Sqeak() = %q , want : %q", got, want)
	}
}
