package list

import (
	"slices"
	"testing"
)

func TestFn(t *testing.T) {
	// Filter, CountFunc
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	fn := func(x int) bool { return x%2 == 0 }
	want, wantCount := []int{2, 4, 6}, 3
	actual := Filter(numbers, fn)
	if slices.Equal(want, actual) == false {
		t.Errorf("Filter() = %v, want %v", actual, want)
	}
	actualCount := CountFunc(numbers, fn)
	if actualCount != wantCount {
		t.Errorf("CountFunc() = %v, want %v", actualCount, wantCount)
	}
	fn = func(x int) bool { return x > 10 }
	want, wantCount = []int{}, 0
	actual = Filter(numbers, fn)
	if slices.Equal(want, actual) == false {
		t.Errorf("Filter() = %v, want %v", actual, want)
	}
	actualCount = CountFunc(numbers, fn)
	if actualCount != wantCount {
		t.Errorf("CountFunc() = %v, want %v", actualCount, wantCount)
	}
	fn = func(x int) bool { return x <= 10 }
	want, wantCount = numbers, len(numbers)
	actual = Filter(numbers, fn)
	if slices.Equal(want, actual) == false {
		t.Errorf("Filter() = %v, want %v", actual, want)
	}
	actualCount = CountFunc(numbers, fn)
	if actualCount != wantCount {
		t.Errorf("CountFunc() = %v, want %v", actualCount, wantCount)
	}
	// FilterIndexed
	want = []int{1, 2, 4, 6, 7}
	actual = FilterIndexed(numbers, func(i, x int) bool { return x%2 == 0 || i%3 == 0 })
	if slices.Equal(want, actual) == false {
		t.Errorf("FilterIndexed() = %v, want %v", actual, want)
	}
	// Reduce
	wantSum := 28
	actualSum := Reduce(numbers, 0, func(result, item int) int {
		return result + item
	})
	if wantSum != actualSum {
		t.Errorf("Reduce() = %d, want %d", actualSum, wantSum)
	}
	// Apply
	want = []int{2, 4, 6, 8, 10, 12, 14}
	actual = Apply(numbers, func(x int) int { return x * 2 })
	if slices.Equal(want, actual) == false {
		t.Errorf("Apply() = %v, want %v", actual, want)
	}
}

func TestFnMap(t *testing.T) {
	// TODO: Map, MapIndexed
	// TODO: MapIf, MapIndexedIf
	// TODO: MapList, MapLookup
}

func TestSumProduct(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	// Sum
	actual, want := Sum(items), 21
	if actual != want {
		t.Errorf("Sum() = %d, want %d", actual, want)
	}
	// Product
	actual, want = Product(items), 720
	if actual != want {
		t.Errorf("Product() = %d, want %d", actual, want)
	}
}
