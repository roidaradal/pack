package str

import "testing"

func TestBuilder(t *testing.T) {
	b := NewBuilder()
	b.Add("1")
	b.AddItems("2", "3", "4")
	b.AddFmt("%d,%d", 5, 6)
	actual := b.Build(",")
	want := "1,2,3,4,5,6"
	if actual != want {
		t.Errorf("Builder: got %q, want %q", actual, want)
	}
}

func TestRepeat(t *testing.T) {
	type testCase struct {
		want       string
		count      int
		text, glue string
	}
	testCases := []testCase{
		{"aaaaa", 5, "a", ""},
		{"ab-ab-ab", 3, "ab", "-"},
		{"", 0, "a", "x"},
		{"b,b", 2, "b", ","},
		{"x", 1, "x", "x"},
	}
	for _, x := range testCases {
		actual := Repeat(x.count, x.text, x.glue)
		if actual != x.want {
			t.Errorf("Repeat: got %q, want %q", actual, x.want)
		}
	}
}
