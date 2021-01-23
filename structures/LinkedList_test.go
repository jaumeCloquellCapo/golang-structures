package structures

import (
	"testing"
)

func TestSinglyLinked_GetFirstValue(t *testing.T) {
	list := NewLinkedList()

	_, err := list.GetFirstValue()
	if err == nil {
		t.Errorf(err.Error())
	}


	list.Add("abc")

	value, err := list.GetFirstValue()
	if err != nil {
		t.Errorf(err.Error())
	}

	if value != "abc" {
		t.Errorf("Expected abc get %s", value)
	}

	list.Add("def")

	value, err = list.GetFirstValue()
	if err != nil {
		t.Errorf(err.Error())
	}

	if value != "abc" {
		t.Errorf("Expected def get %s", value)
	}


	list.InsertAt(0, "xyz")

	value, err = list.GetFirstValue()
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != "xyz" {
		t.Errorf("Expected xyz get %s", value)
	}

}
