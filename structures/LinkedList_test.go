package structures

import (
	"testing"
)

func TestLinkedList_CreateLinkedList(t *testing.T) {
	list := NewLinkedList()

	_, err := list.GetFirstValue()

	if err == nil {
		t.Errorf(err.Error())
	}

}


func TestLinkedList_GetFirstValue(t *testing.T) {
	list := NewLinkedList()

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

func TestLinkedList_GetLastValue(t *testing.T) {
	list := NewLinkedList()

	list.Add("abc")

	value, err := list.GetLastValue()

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != "abc" {
		t.Errorf("Expected def abc %s", value)
	}

	list.Add("def")

	value, err = list.GetLastValue()
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != "def" {
		t.Errorf("Expected def get %s", value)
	}

	list.InsertAt(2, "xyz")

	value, err = list.GetLastValue()
	if err != nil {
		t.Errorf(err.Error())
	}
}

