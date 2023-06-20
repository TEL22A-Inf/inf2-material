package textanalyse

import (
	"reflect"
	"testing"
)

func TestSearchSubstring(t *testing.T) {
	// Testdaten
	text := "abc def abcd deg"

	// Testfall: Substring, der an zwei Stellen vorkommt, einmal ganz am Anfang
	s1 := "abc"
	result1 := SearchSubstring(text, s1)
	expect1 := []int{0, 8}
	if !reflect.DeepEqual(expect1, result1) {
		t.Errorf("result doesn't match expected list. result: %v, expected %v", result1, expect1)
	}

	// Testfall: Substring, der an zwei Stellen vorkommt, aber nicht am Anfang
	s2 := "de"
	result2 := SearchSubstring(text, s2)
	expect2 := []int{4, 13}
	if !reflect.DeepEqual(expect2, result2) {
		t.Errorf("result doesn't match expected list. result: %v, expected %v", result2, expect2)
	}

	// Testfall: Substring, von dem es unterschiedliche Fortsetzungen gibt, die alle vorkommen.

	// Testfall: Substring, der nicht vorkommt.
}
