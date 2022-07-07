package main

import "testing"

func TestMulNumbers(t *testing.T) {
	v := TwoNumbers{3, 4}

	got := v.MulNumbers()
	want := 12

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
