package container_test

import (
	"log"
	"testing"

	. "github.com/febrian-430/data-structure-doodles/container"
)

func TestNodeNext(t *testing.T) {
	tests := []struct {
		context   string
		val       []int
		numOfCall int
		expectVal int
		expectNil bool
	}{
		{context: "when next is not nil", val: []int{1, 2}, numOfCall: 1, expectVal: 2},
		{context: "when linked list has two elements and call next three times", val: []int{1, 2}, numOfCall: 2, expectNil: true},
	}

	for _, testcase := range tests {
		log.Print(testcase.context)
		list := NewLinkedList()
		list.Push(testcase.val...)

		got := list.Head()
		for i := 0; i < testcase.numOfCall; i++ {
			got = got.Next()
		}

		if (got == nil) != testcase.expectNil {
			t.Logf("Expected nil, got %v", got)
			t.Fail()
		}

		if !testcase.expectNil && testcase.expectVal != got.Val {
			t.Logf("Expected %v, got %v", testcase.expectVal, got.Val)
			t.Fail()
		}
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		context   string
		val       []int
		numOfCall int
		expectVal int
		expectNil bool
	}{
		{context: "when has empty list and call next once", val: []int{}, numOfCall: 1, expectNil: true},
		{context: "when has one element and call next once", val: []int{1}, numOfCall: 1, expectVal: 1},
		{context: "when has two elements and call next twice", val: []int{1, 2}, numOfCall: 2, expectVal: 2},
		{context: "when linked list has two elements and call next three times", val: []int{1, 2}, numOfCall: 3, expectNil: true},
	}

	for _, testcase := range tests {
		list := NewLinkedList()
		list.Push(testcase.val...)
		var got *LinkedListNode

		for i := 0; i < testcase.numOfCall; i++ {
			got = list.Next()
		}

		if (got == nil) != testcase.expectNil {
			t.Log(testcase.context)

			t.Logf("Expected nil, got %v", got)
			t.Fail()
		}

		if !testcase.expectNil && testcase.expectVal != got.Val {
			t.Log(testcase.context)

			t.Logf("Expected %v, got %v", testcase.expectVal, got.Val)
			t.Fail()
		}
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		context     string
		val         []int
		expectError bool
	}{
		{context: "should insert normally", val: []int{1, 2}, expectError: false},
		{context: "should fail because the last '1' is lower than tail", val: []int{1, 2, 1}, expectError: true},
		{context: "should fail because tail is higher than four", val: []int{1, 6, 4}, expectError: true},
		{context: "should insert because exclusive", val: []int{1, 6, 5}, expectError: false},
	}

	for _, testcase := range tests {
		list := NewLinkedList()
		err := list.Push(testcase.val...)
		hasError := (err != nil)
		log.Printf("hasError %v", hasError)
		if hasError != testcase.expectError {
			t.Logf("On input: %v, \nExpected error %v, got %v", testcase.val, testcase.expectError, hasError)
			t.Fail()
		}
	}

}
