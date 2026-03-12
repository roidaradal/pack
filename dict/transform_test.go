package dict

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"testing"
)

func TestZip(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := []int{1, 2, 3}
	m := Zip(keys, values)
	for i, key := range keys {
		want := values[i]
		actual, ok := m[key]
		if actual != want || !ok {
			t.Errorf("Zip[%q] = %d, %v, want %d, true", key, actual, ok, want)
		}
		key = key + "x"
		actual, ok = m[key]
		if actual != 0 || ok {
			t.Errorf("Zip[%q] = %d, %v, want 0, false", key, actual, ok)
		}
	}
	keys2, values2 := Unzip(m)
	m2 := Zip(keys2, values2)
	if maps.Equal(m, m2) == false {
		t.Errorf("Unzip.Zip = %v, want %v", m2, m)
	}
	values2 = []int{1, 2}
	m3 := Zip(keys, values2)
	if Len(m3) != 2 {
		t.Errorf("ZipMap.Len = %d, want 2", Len(m3))
	}
}

func TestSwap(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	want := []Entry[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}
	m2 := Swap(m)
	actual := SortedEntries(m2)
	if slices.Equal(actual, want) == false {
		t.Errorf("Swap.Entries = %v, want %v", actual, want)
	}
}

func TestSwapList(t *testing.T) {
	m := map[string][]int{
		"a": {1, 3, 5},
		"b": {2, 4},
	}
	want := []Entry[int, string]{{1, "a"}, {2, "b"}, {3, "a"}, {4, "b"}, {5, "a"}}
	m2 := SwapList(m)
	actual := SortedEntries(m2)
	if slices.Equal(actual, want) == false {
		t.Errorf("SwapList.Entries = %v, want %v", actual, want)
	}
}

func TestFromStruct(t *testing.T) {
	type config struct {
		A, B, C int
	}
	cfg := &config{1, 2, 3}
	want := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	actual, err := FromStruct[int](cfg)
	if err != nil {
		t.Errorf("FromStruct error: %s", err.Error())
	}
	if maps.Equal(actual, want) == false {
		t.Errorf("FromStruct = %v, want %v", actual, want)
	}
	// Test nil input
	want = map[string]int{}
	actual, err = FromStruct[int, config](nil)
	if err != nil {
		t.Errorf("FromStruct error: %s", err.Error())
	}
	if maps.Equal(actual, want) == false {
		t.Errorf("FromStruct = %v, want %v", actual, want)
	}
	// Test unmarshal error
	actual2, err := FromStruct[string, config](cfg)
	if actual2 != nil || err == nil {
		t.Errorf("FromStruct = %v, %v, want nil, error", actual2, err)
	}
	// Test marshal error
	type config2 struct {
		Item any
	}
	cfg2 := &config2{Item: make(chan int)}
	actual3, err := FromStruct[any, config2](cfg2)
	if actual3 != nil || err == nil {
		t.Errorf("FromStruct = %v, %v, want nil, error", actual3, err)
	}
}

func TestToStruct(t *testing.T) {
	type config struct {
		A, B, C int
	}
	obj := Object{"A": 1, "B": 2, "C": 3}
	want := &config{1, 2, 3}
	actual, err := ToStruct[config](obj)
	if err != nil {
		t.Errorf("ToStruct error: %s", err.Error())
	}
	if reflect.DeepEqual(actual, want) == false {
		t.Errorf("ToStruct = %v, want %v", actual, want)
	}
	// Test nil input
	want = &config{0, 0, 0}
	actual, err = ToStruct[config](nil)
	if err != nil {
		t.Errorf("ToStruct error: %s", err.Error())
	}
	if reflect.DeepEqual(actual, want) == false {
		t.Errorf("ToStruct = %v, want %v", actual, want)
	}
	type config2 struct {
		A, B string
	}
	// Test unmarshal error
	actual2, err := ToStruct[config2](obj)
	if actual2 != nil || err == nil {
		t.Errorf("ToStruct = %v, %v, want nil, error", actual2, err)
	}
	// Test marshal error
	obj2 := Object{"A": make(chan int), "B": 5}
	actual3, err := ToStruct[config2](obj2)
	if actual3 != nil || err == nil {
		t.Errorf("ToStruct = %v, %v, want nil, error", actual3, err)
	}
}

func TestToObject(t *testing.T) {
	type config struct {
		A, B, C int
	}
	cfg := &config{1, 2, 3}
	want := Object{"A": 1, "B": 2, "C": 3}
	actual, err := ToObject(cfg)
	if err != nil {
		t.Errorf("ToObject error: %s", err.Error())
	}
	compareObjects(t, actual, want)
}

func TestPruned(t *testing.T) {
	type config struct {
		A, B, C int
	}
	cfg := &config{1, 2, 3}
	want := Object{"B": 2, "C": 3}
	actual, err := Pruned(cfg, "B", "C")
	if err != nil {
		t.Errorf("Pruned error: %s", err.Error())
	}
	compareObjects(t, actual, want)

	type config2 struct {
		Item any
	}
	cfg2 := &config2{Item: make(chan int)}
	actual2, err := Pruned(cfg2, "Item")
	if actual2 != nil || err == nil {
		t.Errorf("Pruned = %v, %v, want nil, error", actual2, err)
	}
}

func compareObjects(t *testing.T, actual, want Object) {
	// Note: cannot use maps.Equal because map value type is <any>
	// The <any> type is not comparable, so even if the map values are the same, the comparison fails
	actualKeys, wantKeys := SortedKeys(actual), SortedKeys(want)
	if slices.Equal(actualKeys, wantKeys) == false {
		t.Errorf("ToObject Keys = %v, want %v", actualKeys, wantKeys)
	}
	for _, key := range wantKeys {
		wantValue := fmt.Sprintf("%v", want[key])
		actualValue := fmt.Sprintf("%v", actual[key])
		if wantValue != actualValue {
			t.Errorf("ToObject[%q] = %s, want %s", key, actualValue, wantValue)
		}
	}
}
