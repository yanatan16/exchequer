package exchequer

import (
	"fmt"
	"reflect"
)

type I interface{}
type A []interface{}
type M map[string]interface{}

type PathDoesntExist string
type TypeCastIsntValid string

func NewPathDoesntExist(x interface{}) PathDoesntExist {
	return PathDoesntExist(fmt.Sprintf("%v", x))
}

func NewTypeCastIsntValid(x interface{}, cast string) TypeCastIsntValid {
	return TypeCastIsntValid(fmt.Sprintf("%v to %s (it is %v)", x, cast, reflect.TypeOf(x)))
}

func (path PathDoesntExist) Error() string {
	return "Path doesn't exist: " + string(path)
}

func (typ TypeCastIsntValid) Error() string {
	return "Type-cast isn't valid: " + string(typ)
}


var mapType reflect.Type = reflect.TypeOf(map[string]interface{}{})
func convertToMap(i interface{}) (map[string]interface{}, bool) {
	// See if it is a map first
	if m, ok := i.(map[string]interface{}); ok {
		return m, true
	} else if i == nil {
		return nil, false
	}

	// Now try converting it
	v := reflect.ValueOf(i)
	if v.Type().ConvertibleTo(mapType) {
		return v.Convert(mapType).Interface().(map[string]interface{}), true
	}

	return nil, false
}

var arrayType reflect.Type = reflect.TypeOf([]interface{}{})
func convertToArray(i interface{}) ([]interface{}, bool) {
	// See if it is an array first
	if a, ok := i.([]interface{}); ok {
		return a, true
	} else if i == nil {
		return nil, false
	}

	// Now try converting it
	v := reflect.ValueOf(i)
	if v.Type().ConvertibleTo(arrayType) {
		return v.Convert(arrayType).Interface().([]interface{}), true
	}

	return nil, false
}

func Get(i I, paths ...interface{}) (interface{}, error) {
	for _, path := range paths {
		if s, ok := path.(string); ok {
			if m, ok := convertToMap(i); ok {
				i = m[s]
				continue
			}
			return nil, NewPathDoesntExist(path)
		}
		if x, ok := path.(int); ok {
			if a, ok := convertToArray(i); ok {
				if x < 0 {
					x = len(a) + x
				}

				if x < 0 || x >= len(a) {
					return nil, NewPathDoesntExist(path)
				} else {
					i = a[x]
					continue
				}
			} else {
				return nil, NewPathDoesntExist(path)
			}
		}
	}

	return i, nil
}

func Set(i I, value interface{}, paths ...interface{}) error {
	for j, path := range paths {
		if s, ok := path.(string); ok {
			if m, ok := convertToMap(i); ok {
				if j < len(paths)-1 {
					if _, ok = m[s]; !ok {
						m[s] = make(map[string]interface{})
					}
					i = m[s]
					continue
				} else {
					m[s] = value
					return nil
				}
			} else {
				return NewPathDoesntExist(path)
			}
		}
		if x, ok := path.(int); ok {
			if a, ok := convertToArray(i); ok {
				if x < 0 {
					x = len(a) + x
				}

				if x < 0 || x >= len(a) {
					return NewPathDoesntExist(path)
				} else {
					if j < len(paths)-1 {
						i = a[x]
						continue
					} else {
						a[x] = value
						return nil
					}
				}
			} else {
				return NewPathDoesntExist(path)
			}
		}
	}
	return nil
}

func String(i I, paths ...interface{}) (string, error) {
	i, err := Get(i, paths...)
	if err != nil {
		return "", err
	}

	if s, ok := i.(string); ok {
		return s, nil
	}

	return "", NewTypeCastIsntValid(i, "string")
}

func Int(i I, paths ...interface{}) (int, error) {
	i, err := Get(i, paths...)
	if err != nil {
		return 0, err
	}

	if s, ok := i.(int); ok {
		return s, nil
	}

	return 0, NewTypeCastIsntValid(i, "int")
}

func Bool(i I, paths ...interface{}) (bool, error) {
	i, err := Get(i, paths...)
	if err != nil {
		return false, err
	}

	if s, ok := i.(bool); ok {
		return s, nil
	}

	return false, NewTypeCastIsntValid(i, "bool")
}

func Float(i I, paths ...interface{}) (float64, error) {
	i, err := Get(i, paths...)
	if err != nil {
		return 0, err
	}

	if s, ok := i.(float64); ok {
		return s, nil
	}

	return 0, NewTypeCastIsntValid(i, "float")
}

func Map(i I, paths ...interface{}) (M, error) {
	i, err := Get(i, paths...)
	if err != nil {
		return nil, err
	}

	if s, ok := convertToMap(i); ok {
		return M(s), nil
	}

	return nil, NewTypeCastIsntValid(i, "map")
}

func Array(i I, paths ...interface{}) (A, error) {
	i, err := Get(i, paths...)
	if err != nil {
		return nil, err
	}

	if s, ok := convertToArray(i); ok {
		return A(s), nil
	}

	return nil, NewTypeCastIsntValid(i, "array")
}

type Q struct {
	i I
	prefix []interface{}
}

func New(i I, prefix_paths...interface{}) *Q {
	return &Q{i, prefix_paths}
}
func (q *Q) Prefix(paths ...interface{}) *Q {
	return New(q.i, append(q.prefix, paths...)...)
}
func (q *Q) Unprefix() *Q {
	return New(q.i)
}
func (q *Q) I() interface{} {
	return q.i
}
func (q *Q) Q(paths ...interface{}) (*Q, error) {
	if i, err := Get(q.i, append(q.prefix, paths...)...); err != nil {
		return nil, err
	} else {
		return New(i), nil
	}
}
func (q *Q) Get(paths ...interface{}) (interface{}, error) {
	return Get(q.i, append(q.prefix, paths...)...)
}
func (q *Q) String(paths ...interface{}) (string, error) {
	return String(q.i, append(q.prefix, paths...)...)
}
func (q *Q) Int(paths ...interface{}) (int, error) {
	return Int(q.i, append(q.prefix, paths...)...)
}
func (q *Q) Bool(paths ...interface{}) (bool, error) {
	return Bool(q.i, append(q.prefix, paths...)...)
}
func (q *Q) Float(paths ...interface{}) (float64, error) {
	return Float(q.i, append(q.prefix, paths...)...)
}
func (q *Q) Map(paths ...interface{}) (M, error) {
	return Map(q.i, append(q.prefix, paths...)...)
}
func (q *Q) Array(paths ...interface{}) (A, error) {
	return Array(q.i, append(q.prefix, paths...)...)
}
func (q *Q) Set(value interface{}, paths ...interface{}) error {
	return Set(q.i, value, append(q.prefix, paths...)...)
}