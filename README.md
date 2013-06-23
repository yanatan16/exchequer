# exchequer [![Build Status][1]][2]

A `map[string]interface{}` query utility belt. I didn't like many of the [other](https://github.com/jmoiron/jsonq) [query](http://godoc.org/github.com/bitly/go-simplejson) libraries for arbitrary objects in Go, so I made my own. Specifically, I wanted one-liner capability with reflection of array index or map key without tying it directly to JSON.

## Documentation

Examples below. [API Documentation on Godoc](http://godoc.org/github.com/yanatan16/exchequer)

## Example

Suppose we have a arbitrary object in Go, perhaps derived from the following JSON:

```
{
	"foo": "bar",
	"baz": 123,
	"mux": {
		"flux": "capaciter",
		"marry-me": false,
		"fifty": {
			"cents": [1,2,3,4,5,6,7,8,9,10]
		}
	},
	"shifty": {
		"one": 1,
		"two": 2,
		"five": 5.55
	}
}
```

This is what you are doing now:

```go
func getFooAsString(object interface{}) (string, error) {
	// You could do this...
	lvl1, ok := object.(map[string]interface{})
	if !ok {
		return nil, errors.New("wah!")
	}

	lvl2, ok := lvl1["foo"]
	if !ok {
		return nil, errors.New("please")
	}

	if foo, ok := lvl2.(string); ok {
		return foo, nil
	}

	return nil, errors.New("help!")
}
```

With exchequer, you can make this much easier:

```go
import exq "exchequer"

// obj.foo as string
exq.String(obj, "foo")

// obj.baz as int
exq.Int(obj, "baz")

// obj.mux.fifty.cents[5] as int
exq.Int(obj, "mux", "fifty", "cents", 5)

// Create a query-able object
q := exq.New(obj)

// obj.shifty as a map
q.Map("shifty")

// obj.mux["marry-me"] as bool
q.Bool("mux", "marry-me")
```

## License

MIT found in LICENSE file



[1]: https://travis-ci.org/yanatan16/exchequer.png?branch=master
[2]: http://travis-ci.org/yanatan16/exchequer