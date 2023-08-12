package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Soma(15, 15) = %d; want 30", total)
	}
}

func TestSub(t *testing.T) {
	total := Sub(50, 40)
	if total != 10 {
		t.Errorf("Sub(50, 40) = %d; want 10", total)
	}
}

func TestTimes(t *testing.T) {
	total := Times(2, 2)
	if total != 4 {
		t.Errorf("Sub(2, 2) = %d; want 4", total)
	}
}
