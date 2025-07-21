package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5.0
	if result != expected {
		t.Errorf("Add(2, 3) = %.2f; want %.2f", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := Sub(10, 4)
	expected := 6.0
	if result != expected {
		t.Errorf("Sub(10, 4) = %.2f; want %.2f", result, expected)
	}
}