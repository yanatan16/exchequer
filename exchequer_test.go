package exchequer

import (
	"testing"
)

var testobj interface{} = map[string]interface{}{
	"foo": "bar",
	"baz": 123,
	"mux": map[string]interface{}{
		"flux":     "capaciter",
		"marry-me": false,
		"fifty": map[string]interface{}{
			"cents": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	},
	"shifty": map[string]interface{}{
		"one":  1,
		"two":  2,
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
	if i, err := Int(testobj, "mux", "fifty", "cents", 5); err != nil {
		t.Error(err)
	} else if i != 6 {
		t.Error("mux.fifty.cents.5 != 6", i)
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
			if x != i+1 {
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

func TestObjectString(t *testing.T) {
	q := New(testobj)
	if s, err := q.String("foo"); err != nil {
		t.Error(err)
	} else if s != "bar" {
		t.Error("foo != bar", s)
	}
}

func TestObjectInt(t *testing.T) {
	q := New(testobj)
	if i, err := q.Int("mux", "fifty", "cents", 5); err != nil {
		t.Error(err)
	} else if i != 6 {
		t.Error("mux.fifty.cents.5 != 6", i)
	}
}

func TestObjectBool(t *testing.T) {
	q := New(testobj)
	if i, err := q.Bool("mux", "marry-me"); err != nil {
		t.Error(err)
	} else if i != false {
		t.Error("mux.marry-me != false", i)
	}
}

func TestObjectFloat(t *testing.T) {
	q := New(testobj)
	if i, err := q.Float("shifty", "five"); err != nil {
		t.Error(err)
	} else if i != 5.55 {
		t.Error("shifty.five != 5.55", i)
	}
}

func TestObjectArray(t *testing.T) {
	q := New(testobj)
	if arr, err := q.Array("mux", "fifty", "cents"); err != nil {
		t.Error(err)
	} else {
		for i, x := range arr {
			if x != i+1 {
				t.Error("mux.fifty.cents != range(1,11)", arr)
				break
			}
		}
	}
}

func TestObjectMap(t *testing.T) {
	q := New(testobj)
	if obj, err := q.Map("shifty"); err != nil {
		t.Error(err)
	} else {
		if len(obj) != 3 || obj["one"] != 1 || obj["two"] != 2 || obj["five"] != 5.55 {
			t.Error("shifty != map[one:1...]", obj)
		}
	}
}

func TestSet(t *testing.T) {
	i := I(testobj)
	Set(i, "herro", "hello", "konichiwa", "mygod")
	if v, err := String(i, "hello", "konichiwa", "mygod"); err != nil {
		t.Error(err)
	} else if v != "herro" {
		t.Error("v isn't herro " + v)
	}

	Set(i, 10, "mux", "fifty", "cents", 0)
	if vi, ierr := Int(i, "mux", "fifty", "cents", 0); ierr != nil {
		t.Error(ierr)
	} else if vi != 10 {
		t.Error("vi isn't 10 " + string(vi))
	}
}
