package lang

import "testing"

func TestIdentity(t *testing.T) {
	fn1 := Identity[int]
	actual, want := fn1(5), 5
	if actual != want {
		t.Errorf("Identity(%d) = %d, want %d", want, actual, want)
	}
	fn2 := Identity[string]
	actual2, want2 := fn2("a"), "a"
	if actual2 != want2 {
		t.Errorf("Identity(%q) = %q, want %q", want2, actual2, want2)
	}
}
