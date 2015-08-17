package solution

import "testing"

func TestPushInt(t *testing.T) {
	input := "1"
	actual := Solution(input)
	expected := 1

	if expected != actual {
		t.Errorf("Should be 1, was: ", actual)
	}
}

func TestPush2Ints(t *testing.T) {
	input := "23"
	actual := Solution(input)
	expected := 3

	if expected != actual {
		t.Errorf("Should be 2, was: ", actual)
	}
}

func TestAdd2ints(t *testing.T) {
	input := "23+"
	actual := Solution(input)
	expected := 5

	if expected != actual {
		t.Errorf("Should be %d, was: %d", expected, actual)
	}
}

func TestMultiply2ints(t *testing.T) {
	input := "23*"
	actual := Solution(input)
	expected := 6

	if expected != actual {
		t.Errorf("Should be %d, was: %d", expected, actual)
	}
}

func TestAddAndMultiple(t *testing.T) {
	input := "32+4*3+2*"
	actual := Solution(input)
	expected := 46

	if expected != actual {
		t.Errorf("Should be %d, was: %d", expected, actual)
	}
}
