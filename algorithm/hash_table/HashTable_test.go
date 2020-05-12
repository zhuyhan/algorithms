package hash_table

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	h := New(100)
	h.Put("go", "hello world")
	h.Put("php", "good")
	h.Put("java", "better")

	value, err := h.Get("go")
	if err != nil {
		t.Error(err.Error())
	}
	if value != "hello world" {
		t.Errorf("want = hello world,got=%v", value)
	}
	h.Put("go", "hello world!")
	value, err = h.Get("go")
	if err != nil {
		t.Error(err.Error())
	}

	if value != "hello world!" {
		t.Errorf("want = hello world!,got=%v", value)
	}

	err = h.Remove("go")
	if err != nil {
		t.Error(err.Error())
	}
	value, err = h.Get("go")
	if value != "" || err == nil || h.size != 2 {
		t.Error()
	}

}
