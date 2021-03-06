package array

import (
	"testing"
)

func TestNew(t *testing.T) {
	arr := New(0)
	if arr != nil {
		t.Fatal("New zero value Array did not work as expected.")
	}

	arr = New(3)
	if arr == nil {
		t.Fatal("New did not work as expected.")
	}
}

func TestGet(t *testing.T) {
	arr := New(3)
	v, _ := arr.Get(1)
	if v != 0 {
		t.Fatal("Get did not work as expected.")
	}

	_, err := arr.Get(3)
	if err == nil || err.Error() != "Out of index range" {
		t.Fatal("Get did not work as expected.")
	}
}

func TestInsert(t *testing.T) {
	arr := New(3)
	err := arr.Insert(0, 1)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}
	v, _ := arr.Get(0)
	if v != 1 {
		t.Fatal("Insert did not work as expected.")
	}

	err = arr.Insert(3, 1)
	if err == nil || err.Error() != "Out of index range" {
		t.Fatal("Insert did not work as expected.")
	}

	err = arr.Insert(0, 2)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}
	v, _ = arr.Get(1)
	if v != 1 {
		t.Fatal("Insert did not work as expected.")
	}

	err = arr.Insert(0, 3)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}
	err = arr.Insert(0, 4)
	if err == nil || err.Error() != "Array is full" {
		t.Fatal("Insert did not work as expected.")
	}
	v, _ = arr.Get(2)
	if v != 1 {
		t.Fatal("Insert did not work as expected.")
	}
}

func TestInsertEnding(t *testing.T) {
	arr := New(3)
	err := arr.InsertEnding(1)
	if err != nil {
		t.Fatal("InsertEnding did not work as expected.")
	}
	err = arr.InsertEnding(3)
	if err != nil {
		t.Fatal("InsertEnding did not work as expected.")
	}
	err = arr.InsertEnding(2)
	if err != nil {
		t.Fatal("InsertEnding did not work as expected.")
	}

	if arr.print() != "1|3|2|" {
		t.Fatal("InsertEnding did not work as expected.")
	}
}

func TestDelete(t *testing.T) {
	arr := New(3)
	err := arr.Insert(0, 10)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}

	err = arr.Delete(0)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}

	v, _ := arr.Get(0)
	if v != 0 {
		t.Fatal("Delete did not work as expected.")
	}

	err = arr.Delete(1)
	if err == nil || err.Error() != "Array is empty" {
		t.Fatal("Delete did not work as expected.")
	}

	err = arr.Delete(10)
	if err == nil || err.Error() != "Out of index range" {
		t.Fatal("Delete did not work as expected.")
	}

	err = arr.Insert(0, 10)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}
	err = arr.Insert(0, 20)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}
	err = arr.Insert(0, 30)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}

	err = arr.Delete(1)
	if err != nil {
		t.Fatal("Delete did not work as expected.")
	}

	v, _ = arr.Get(0)
	if v != 30 {
		t.Fatal("Delete did not work as expected.")
	}
	v, _ = arr.Get(2)
	if v != 10 {
		t.Fatal("Delete did not work as expected.")
	}
}

func TestLen(t *testing.T) {
	arr := New(3)
	l := arr.Len()
	if l != 0 {
		t.Fatal("Len did not work as expected.")
	}

	err := arr.Insert(0, 1)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}
	err = arr.Insert(0, 2)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}

	l = arr.Len()
	if l != 2 {
		t.Fatal("Len did not work as expected.")
	}
}
