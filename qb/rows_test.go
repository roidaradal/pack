package qb

import (
	"fmt"
	"maps"
	"reflect"
	"testing"

	"github.com/roidaradal/pack/dict"
	"github.com/roidaradal/pack/dyn"
)

type mockScanner struct {
	items []any
}

func (m mockScanner) Scan(fieldRefs ...any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic encountered: %v", r)
		}
	}()
	if len(fieldRefs) != len(m.items) {
		return fmt.Errorf("expected %d fieldRefs, got %d", len(m.items), len(fieldRefs))
	}
	for i, fieldRef := range fieldRefs {
		fieldValue := dyn.MustDerefValue(fieldRef)
		fieldValue.Set(reflect.ValueOf(m.items[i]))
	}
	return err
}

func TestRowFunctions(t *testing.T) {
	type User struct {
		Name     string
		Password string
		Age      int
		secret   string
	}
	type School struct {
		Name    string
		Address string
	}
	this := NewInstance(MySQL)
	user := &User{"john", "123456", 25, "secret"}
	school := &School{"UP", "Lahug"}
	userRef := new(User)
	err := AddType(this, userRef)
	if err != nil {
		t.Errorf("AddType() error = %v", err)
	}
	// ToRow
	empty := dict.Object{}
	userObj := dict.Object{"Name": "john", "Password": "123456", "Age": 25}
	testCases := [][2]dict.Object{
		{userObj, ToRow(this, user)},
		{empty, ToRow(this, school)},
	}
	for _, x := range testCases {
		want, actual := x[0], x[1]
		if maps.Equal(want, actual) == false {
			t.Errorf("ToRow() = %v, want %v", actual, want)
		}
	}
	// Not a struct type
	intReader := NewRowReader[int](this, "Value", "Decimal")
	intResult := intReader(mockScanner{})
	if intResult.IsError() == false || intResult.Value() != 0 {
		t.Errorf("NewRowReader[int] should return an error")
	}
	// Valid full reader
	fullReader := FullRowReader(this, userRef)
	if fullReader == nil {
		t.Errorf("FullRowReader() should return a rowReader, got nil")
	}
	// Successful read
	result := fullReader(mockScanner{items: []any{"John", "111", 20}})
	if result.NotError() == false {
		t.Errorf("FullRowReader() read = %v, want non-error", result)
	}
	// Check that struct has been filled after fullReader read
	want := User{"John", "111", 20, ""}
	if want != result.Value() {
		t.Errorf("FullRowReader() read = %v, want %v", result.Value(), want)
	}
	// Valid row reader, with specified columns
	nameCol, pwdCol := this.Column(&userRef.Name), this.Column(&userRef.Password)
	rowReader := NewRowReader[User](this, nameCol, pwdCol)
	result = rowReader(mockScanner{items: []any{"Jane", "222"}})
	if result.NotError() == false {
		t.Errorf("RowReader() read = %v, want non-error", result)
	}
	// Check that struct has been filled after rowReader read
	want = User{"Jane", "222", 0, ""}
	if want != result.Value() {
		t.Errorf("RowReader() read = %v, want %v", result.Value(), want)
	}
	emptyUser := User{}
	// Valid row reader, but error in scanning (mocked by incomplete items / invalid type)
	result = rowReader(mockScanner{items: []any{"Jane", 333}})
	if result.IsError() == false || result.Value() != emptyUser {
		t.Errorf("RowReader() read = %v, want Result<%v, error>", result, emptyUser)
	}
	result = rowReader(mockScanner{items: []any{"Jane"}})
	if result.IsError() == false || result.Value() != emptyUser {
		t.Errorf("RowReader() read = %v, want Result<%v, error>", result, emptyUser)
	}
	// Error because of blank columns
	userReader := NewRowReader[User](this, nameCol, pwdCol, "")
	result = userReader(mockScanner{})
	if result.IsError() == false || result.Value() != emptyUser {
		t.Errorf("NewRowReader() read = %v, want Result<%v, error>", result, emptyUser)
	}
	// Error because of unknown column field
	userReader = NewRowReader[User](this, nameCol, pwdCol, "secret")
	result = userReader(mockScanner{})
	if result.IsError() == false || result.Value() != emptyUser {
		t.Errorf("NewRowReader() read = %v, want Result<%v, error>", result, emptyUser)
	}
}
