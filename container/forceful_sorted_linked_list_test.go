package container_test

import (
	"log"
	"testing"

	. "github.com/febrian-430/data-structure-doodles/container"
)

func TestNext(t *testing.T) {

	list := NewLinkedList()
	got := list.Next()
	if nil != got {
		t.Logf("Expected nil, got %v", got)
		t.Fail()
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
		// t.Log(testcase.context)
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
