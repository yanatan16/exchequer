package exchequer

import (
	"testing"
)

var testobj interface{} = map[string]interface{}{
	"foo": "bar",
	"baz": 123,
	"mux": map[string]interface{}{
		"flux": "capaciter",
		"marry-me": false,
		"fifty": map[string]interface{}{
			"cents": []interface{}{1,2,3,4,5,6,7,8,9,10},
		},
	},
	"shifty": map[string]interface{}{
		"one": 1,
		"two": 2,
		"five": 5.55,
	},
}

func TestString(t *testing.T) {
	if s, err := String(testobj, "foo"); err != nil {
		t.Error(err)
	} else if s != "bar" {
		t.Error("foo != bar", s)
	}
}

func TestInt(t *testing.T) {
	if i, err := Int(testobj, "baz"); err != nil {
		t.Error(err)
	} else if i != 123 {
		t.Error("baz != 123", i)
	}
}

func TestBool(t *testing.T) {
	if i, err := Bool(testobj, "mux", "marry-me"); err != nil {
		t.Error(err)
	} else if i != false {
		t.Error("mux.marry-me != false", i)
	}
}

func TestFloat(t *testing.T) {
	if i, err := Float(testobj, "shifty", "five"); err != nil {
		t.Error(err)
	} else if i != 5.55 {
		t.Error("shifty.five != 5.55", i)
	}
}

func TestArray(t *testing.T) {
	if arr, err := Array(testobj, "mux", "fifty", "cents"); err != nil {
		t.Error(err)
	} else {
		for i, x := range arr {
			if x != i + 1 {
				t.Error("mux.fifty.cents != range(1,11)", arr)
				break
			}
		}
	}
}

func TestMap(t *testing.T) {
	if obj, err := Map(testobj, "shifty"); err != nil {
		t.Error(err)
	} else {
		if len(obj) != 3 || obj["one"] != 1 || obj["two"] != 2 || obj["five"] != 5.55 {
			t.Error("shifty != map[one:1...]", obj)
		}
	}
}