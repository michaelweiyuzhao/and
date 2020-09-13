package stack

import (
	"testing"
)

func TestEmptyOnCreate(t *testing.T) {
	myStack := Stack{}
	if !myStack.Empty() {
		t.Errorf("Stack was not empty on creation")
	}
}

func TestPushOnCreateNotEmpty(t *testing.T) {
	myStack := Stack{}
	err := myStack.Push(0)
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

func TestSizeAfterPushes(t *testing.T) {
	myStack := Stack{}
	for i := 0; i < 10; i++ {
		err := myStack.Push(i)
		if err != nil {
			t.Errorf("Stack returned error on push")
		}
	}
	if myStack.Size() != 10 {
		t.Errorf("Stack size is incorrect")
	}
}

func TestSizeAfterPops(t *testing.T) {
	myStack := Stack{}
	for i := 0; i < 10; i++ {
		err := myStack.Push(i)
		if err != nil {
			t.Errorf("Stack returned error on push")
		}
	}
	for i := 0; i < 5; i++ {
		_, err := myStack.Pop()
		if err != nil {
			t.Errorf("Stack returned error on push")
		}
	}
	if myStack.Size() != 5 {
		t.Errorf("Stack size is incorrect")
	}
}

func TestPeekOnCreateCausesError(t *testing.T) {
	myStack := Stack{}
	_, err := myStack.Peek()
	if err == nil {
		t.Errorf("Stack returned no error")
	}
	if !myStack.Empty() {
		t.Errorf("Stack was not empty after pop")
	}
}

func TestPushThenPeekReturnsSameValueAndCorrectSize(t *testing.T) {
	myStack := Stack{}
	expected := 10
	err := myStack.Push(expected)
	if err != nil {
		t.Errorf("Stack returned error on push")
	}
	actual, err := myStack.Peek()
	if err != nil {
		t.Errorf("Stack returned error on pop")
	}
	if actual != expected {
		t.Errorf("Stack returned different value")
	}
	size := myStack.Size()
	if size != 1 {
		t.Errorf("Stack size reduced after peek")
	}
}
