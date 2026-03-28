package dict

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/zeroibot/tst"
)

func TestDict(t *testing.T) {
	// Copy
	m := map[string]int{"apple": 5, "orange": 3, "banana": 2}
	m2 := Copy(m)
	tst.AssertMapEqual(t, "Copy", m2, m)
	// Len
	tst.AssertEqual(t, "Len", Len(m), 3)
	// NotEmpty
	tst.AssertEqual(t, "NotEmpty", NotEmpty(m), true)
	// Keys
	actualKeys := Keys(m)
	slices.Sort(actualKeys)
	tst.AssertListEqual(t, "Keys", actualKeys, []string{"apple", "banana", "orange"})
	// Values
	actualValues := Values(m)
	slices.Sort(actualValues)
	tst.AssertListEqual(t, "Values", actualValues, []int{2, 3, 5})
	// Entries
	wantEntries := []Entry[string, int]{{"apple", 5}, {"banana", 2}, {"orange", 3}}
	actualEntries := Entries(m)
	sortFn := func(e1, e2 Entry[string, int]) int {
		return cmp.Compare(e1.Key, e2.Key)
	}
	slices.SortFunc(actualEntries, sortFn)
	tst.AssertListEqual(t, "Entries", actualEntries, wantEntries)
	// Entry.String
	wantStrings := []string{"<apple: 5>", "<banana: 2>", "<orange: 3>"}
	for i, entry := range actualEntries {
		tst.AssertEqual(t, "Entry.String", entry.String(), wantStrings[i])
	}
	// Entry.Tuple
	want1, want2 := "apple", 5
	actual1, actual2 := wantEntries[0].Tuple()
	tst.AssertEqual(t, "Tuple", actual1, want1)
	tst.AssertEqual(t, "Tuple", actual2, want2)
	// No Key
	noKeyCases := []tst.P2W1[map[string]int, string, bool]{
		{m, "apple", false},
		{m, "grape", true},
	}
	tst.AllP2W1(t, noKeyCases, "NoKey", NoKey, tst.AssertEqual)
	// No Value
	noValueCases := []tst.P2W1[map[string]int, int, bool]{
		{m, 3, false},
		{m, 5, false},
		{m, 1, true},
		{m, 69, true},
	}
	tst.AllP2W1(t, noValueCases, "NoValue", NoValue, tst.AssertEqual)
	// Map Get
	getCases := []tst.P1W1[string, int]{
		{"apple", 5},
		{"zebra", 0},
	}
	for _, x := range getCases {
		key, want := x.P1, x.W1
		tst.AssertEqual(t, fmt.Sprintf("map[%q]", key), m[key], want)
	}
	// GetOrDefault, SetDefault
	defaultValue := 69
	getCases = []tst.P1W1[string, int]{
		{"orange", 3},
		{"cherry", defaultValue},
	}
	for _, x := range getCases {
		key, want := x.P1, x.W1
		tst.AssertEqual(t, "GetOrDefault", GetOrDefault(m, key, defaultValue), want)
		SetDefault(m, key, defaultValue)
		tst.AssertEqual(t, "SetDefault", m[key], want)
	}
	// NoKeyFunc
	keyFnCases := []tst.P2W1[map[string]int, func(string) bool, bool]{
		{m, func(key string) bool { return key == "apple" }, false},
		{m, func(key string) bool { return key == "zebra" }, true},
	}
	tst.AllP2W1(t, keyFnCases, "NoKeyFunc", NoKeyFunc, tst.AssertEqual)
	// No ValueFunc
	valueFnCases := []tst.P2W1[map[string]int, func(int) bool, bool]{
		{m, func(value int) bool { return value > 100 }, true},
		{m, func(value int) bool { return value == 5 }, false},
	}
	tst.AllP2W1(t, valueFnCases, "NoValueFunc", NoValueFunc, tst.AssertEqual)
	// Filter.Entries
	mf := Filter(m, func(key string, value int) bool { return key != "zebra" && value <= 50 })
	tst.AssertListEqual(t, "Filter.Entries", SortedEntries(mf), wantEntries)
	// Update
	m2 = map[string]int{"cherry": 30, "banana": 10}
	wantEntries = []Entry[string, int]{
		{"apple", 5}, {"banana", 10}, {"cherry", 30}, {"orange", 3},
	}
	Update(m, m2)
	tst.AssertListEqual(t, "Update", SortedEntriesFunc(m, sortFn), wantEntries)
	// Clear, IsEmpty
	Clear(m)
	tst.AssertEqual(t, "IsEmpty", IsEmpty(m), true)
}
