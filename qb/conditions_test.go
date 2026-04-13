package qb

import (
	"testing"

	"github.com/zeroibot/pack/list"
	"github.com/zeroibot/tst"
)

func TestConditions(t *testing.T) {
	type Person struct {
		Name    string
		Address string
		Age     int
		Job     string
		Score   int
	}
	p := &Person{}
	this := testPrelude(t, p)

	condNone := NoCondition()
	condEqual := Equal(this, &p.Name, "John")
	condNotEqual := NotEqual(this, &p.Job, "manager")
	condPrefix := Prefix(this, &p.Job, "assistant")
	condSuffix := Suffix(this, &p.Address, "City")
	condSubstring := Substring(this, &p.Address, "Tower")
	condGreater := Greater(this, &p.Score, 75)
	condGreaterEqual := GreaterEqual(this, &p.Age, 20)
	condLesser := Lesser(this, &p.Age, 60)
	condLesserEqual := LesserEqual(this, &p.Score, 50)
	condIn := In(this, &p.Job, []string{"dev", "qa", "intern"})
	condNotIn := NotIn(this, &p.Score, []int{67, 69})
	condAnd := And(condEqual, condGreater)
	condOr := Or(condLesserEqual, condIn)

	testCases := []tst.P1W2[Condition, string, []any]{
		{condNone, "true", []any{}},
		{condEqual, "`Name` = ?", []any{"John"}},
		{condNotEqual, "`Job` <> ?", []any{"manager"}},
		{condPrefix, "`Job` LIKE ?", []any{"assistant%"}},
		{condSuffix, "`Address` LIKE ?", []any{"%City"}},
		{condSubstring, "`Address` LIKE ?", []any{"%Tower%"}},
		{condGreater, "`Score` > ?", []any{75}},
		{condGreaterEqual, "`Age` >= ?", []any{20}},
		{condLesser, "`Age` < ?", []any{60}},
		{condLesserEqual, "`Score` <= ?", []any{50}},
		{condIn, "`Job` IN (?, ?, ?)", []any{"dev", "qa", "intern"}},
		{condNotIn, "`Score` NOT IN (?, ?)", []any{67, 69}},
		{condAnd, "(`Name` = ? AND `Score` > ?)", []any{"John", 75}},
		{condOr, "(`Score` <= ? OR `Job` IN (?, ?, ?))", []any{50, "dev", "qa", "intern"}},
	}
	tst.AllP1W2(t, testCases, "Condition.BuildCondition", Condition.BuildCondition, tst.AssertEqual, tst.AssertListEqual)
}

func TestCombos(t *testing.T) {
	type Person struct {
		Name    string
		Address string
		Age     int
		Job     string
		Score   int
	}
	p := &Person{}
	this := testPrelude(t, p)

	persons := []Person{
		{"John", "Astra Tower, Gotham City", 30, "manager", 95},
		{"Jane", "Star Village, India", 65, "assistant ceo", 85},
		{"Jill", "Tower of Doom, Japan", 20, "dev", 55},
		{"Jim", "Sub Division, Singapore City", 40, "qa", 67},
		{"Jack", "N/A", 18, "intern", 50},
	}

	condNone := NoCondition2[Person]()
	condEqual := Equal2[Person](this, &p.Name, "John")
	condNotEqual := NotEqual2[Person](this, &p.Job, "manager")
	condPrefix := Prefix2[Person](this, &p.Job, "assistant")
	condSuffix := Suffix2[Person](this, &p.Address, "City")
	condSubstring := Substring2[Person](this, &p.Address, "Tower")
	condGreater := Greater2[Person](this, &p.Score, 75)
	condGreaterEqual := GreaterEqual2[Person](this, &p.Age, 20)
	condLesser := Lesser2[Person](this, &p.Age, 60)
	condLesserEqual := LesserEqual2[Person](this, &p.Score, 50)
	condIn := In2[Person](this, &p.Job, []string{"dev", "qa", "intern"})
	condNotIn := NotIn2[Person](this, &p.Score, []int{67, 69})
	condAnd := And2[Person](condEqual, condGreater)
	condOr := Or2[Person](condLesserEqual, condIn)

	testCases := []tst.P1W1[DualCondition[Person], []bool]{
		{condNone, []bool{true, true, true, true, true}},
		{condEqual, []bool{true, false, false, false, false}},
		{condNotEqual, []bool{false, true, true, true, true}},
		{condPrefix, []bool{false, true, false, false, false}},
		{condSuffix, []bool{true, false, false, true, false}},
		{condSubstring, []bool{true, false, true, false, false}},
		{condGreater, []bool{true, true, false, false, false}},
		{condGreaterEqual, []bool{true, true, true, true, false}},
		{condLesser, []bool{true, false, true, true, true}},
		{condLesserEqual, []bool{false, false, false, false, true}},
		{condIn, []bool{false, false, true, true, true}},
		{condNotIn, []bool{true, true, true, false, true}},
		{condAnd, []bool{true, false, false, false, false}},
		{condOr, []bool{false, false, true, true, true}},
	}
	checkTestResults := func(condition DualCondition[Person]) []bool { return list.Map(persons, condition.Test) }
	tst.AllP1W1(t, testCases, "Combo.Test", checkTestResults, tst.AssertListEqual)
}
