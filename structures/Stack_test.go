package structures

import (
	"testing"
)

func TestStackList_Create(t *testing.T) {
	list := NewStack()

	_, err := list.Pull()

	if err == nil {
		t.Errorf(err.Error())
	}

}


func TestStackList_GetLastValue(t *testing.T) {
	list := NewStack()

	list.Push("abc")

	value, err := list.Pull()

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != "abc" {
		t.Errorf("Expected abc get %s", value)
	}

	list.Push("def")

	value, err = list.Pull()

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != "def" {
		t.Errorf("Expected def get %s", value)
	}

	_, err = list.Pull()

	if err == nil {
		t.Errorf(err.Error())
	}

}