package main

import "testing"

func TestAddNumbers(t *testing.T) {
	v := TwoNumbers{3, 4}

	got := v.AddNumbers()
	want := 7

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
