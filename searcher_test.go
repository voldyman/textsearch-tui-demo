package main

import (
	"testing"
)

func TestSearcher(t *testing.T) {
	s := NewSearcher()
	if len(s.Search("cmake")) == 0 {
		t.Fail()
	}
}
