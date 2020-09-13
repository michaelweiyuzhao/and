package stack

import (
	"testing"
	"math/rand"
)

func TestEmptyOnCreate(t *testing.T) {
	myStack := Stack{}
	if !myStack.Empty() {
		t.Errorf("Stack was not empty on creation")
	}
}

func TestPushOnCreateNotEmpty(t *testing.T) {
	myStack := Stack{}
	err := myStack.Push(rand.Int())
	if err != nil {
		t.Errorf("Stack returned error")
	}
	if myStack.Empty() {
		t.Errorf("Stack was empty after push")
	}
}

func TestPopOnCreateCausesError(t *testing.T) {
	myStack := Stack{}
	_, err := myStack.Pop()
	if err == nil {
		t.Errorf("Stack returned no error")
	}
	if !myStack.Empty() {
		t.Errorf("Stack was not empty after pop")
	}
}

func TestEmptyAfterPushThenPop(t *testing.T) {
	myStack := Stack{}
	myStack.Push(10)
	myStack.Pop()
	if !myStack.Empty() {
		t.Errorf("Stack not empty after push and pop")
	}
}

func TestPushThenPopReturnsSameValue(t *testing.T) {
	myStack := Stack{}
	expected := 10
	err := myStack.Push(expected)
	if err != nil {
		t.Errorf("Stack returned error on push")
	}
	actual, err := myStack.Pop()
	if err != nil {
		t.Errorf("Stack returned error on pop")
	}
	if actual != expected {
		t.Errorf("Stack returned different value")
	}
}

func TestPushesAndPopsInStackOrder(t *testing.T) {
	myStack := Stack{}
	for expected := 10; expected < 20; expected++ {
		err := myStack.Push(expected)
		if err != nil {
			t.Errorf("Stack returned error on push")
		}
	}
	for expected := 19; expected >= 10; expected-- {
		actual, err := myStack.Pop()
		if err != nil {
			t.Errorf("Stack returned error on pop")
		}
		if actual != expected {
			t.Errorf("Stack returned different value")
		}
	}
}

func TestEmptyAfterPushesAndPops(t *testing.T) {
	myStack := Stack{}
	myStack.Push(0)
	myStack.Push(0)
	myStack.Pop()
	myStack.Push(0)
	myStack.Pop()
	myStack.Pop()
	myStack.Push(0)
	myStack.Push(0)
	myStack.Pop()
	myStack.Pop()
	if !myStack.Empty() {
		t.Errorf("Stack not empty after push and pop")
	}
}
