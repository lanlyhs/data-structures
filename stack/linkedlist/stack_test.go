package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := NewStack()
	err := stack.Push(1)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}
	if stack.print() != "1|" {
		t.Fatal("Push did not work as expected.")
	}

	err = stack.Push(2)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}

	err = stack.Push(3)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}
	if stack.print() != "1|2|3|" {
		t.Fatal("Push did not work as expected.")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	err := stack.Push(1)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}

	v, err := stack.Pop()
	if err != nil || v != 1 {
		t.Fatal("Pop did not work as expected.")
	}

	_, err = stack.Pop()
	if err == nil || err.Error() != "Stack is empty" {
		t.Fatal("Push did not work as expected.")
	}

	err = stack.Push(2)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}

	err = stack.Push(3)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}

	err = stack.Push(4)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}

	v, err = stack.Pop()
	if err != nil || v != 4 {
		t.Fatal("Pop did not work as expected.")
	}
}
