# Go maps in action

## Introduction

<hr>

One of the most useful data structures in computer science is the hash table. Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes. Go provides a built-in map type that implements a hash table.

<br>

## Declaration and initialization
<hr>

A Go map type looks like this:

```Go
map[KeyType]ValueType
```
where <code>KeyType</code> may be any type that is <code>comparable</code> (more on this later), and <code>ValueType</code> may be any type at all, including another map!

This variable m is a map of string keys to int values:
```Go
var m map[string]int
```

Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function:

```Go
m = make(map[string]int)
```

The <code>make</code> function allocates and initializes a hash map data structure and returns a map value that points to it. The specifics of that data structure are an implementation detail of the runtime and are not specified by the language itself. In this article we will focus on the use of maps, not their implementation.

## Working with maps
<hr>

Go provides a familiar syntax for working with maps. This statement sets the key "route" to the value 66:

```Go
m["route"] = 66
```

This statement retrieves the value stored under the key "route" and assigns it to a new variable i:

```Go
i := m["route"]
```

If the requested key doesn't exist, we get the value type's zero value. In this case the value type is int, so the zero value is 0:

```Go
j := m["root"]
// j == 0
```

The built in <code>len</code> function returns on the number of items in a map:
```Go
n := len(m)
```

The built in <code>delete</code> function removes an entry from the map:
```Go
delete(m, "route")
```

The <code>delete</code> function doesn't return anything, and will do nothing if the specified key doesn't exist.

A two-value assignment tests for the existence of a key:
```Go
i, ok := m["route"]
```

In this statement, the first value (i) is assigned the value stored under the key "route". If that key doesn't exist, i is the value type's zero value (0). The second value (ok) is a bool that is true if the key exists in the map, and false if not.

To test for a key without retrieving the value, use an underscore in place of the first value:

```Go
_, ok := m["route"]
```
To iterate over the contents of a map, use the <code>range</code> keyword:

```Go
for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}
```

To initialize a map with some data, use a map literal:

```Go
commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}
```

The same syntax may be used to initialize an empty map, which is functionally identical to using the make function:

```Go
m = map[string]int{}
```
## Key types
<hr>

As mentioned earlier, map keys may be of any type that is comparable. The language spec defines this precisely, but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.

It's obvious that strings, ints, and other basic types should be available as map keys, but perhaps unexpected are struct keys. Struct can be used to key data by multiple dimensions. For example, this map of maps could be used to tally web page hits by country:
```Go
hits := make(map[string]map[string]int)
```

This is map of string to (map of string to int). Each key of the outer map is the path to a web page with its own inner map. Each inner map key is a two-letter country code. This expression retrieves the number of times an Australian has loaded the documentation page:

```Go
n := hits["/doc/"]["au"]
```

Unfortunately, this approach becomes unwieldy when adding data, as for any given outer key you must check if the inner map exists, and create it if needed:

```Go
func add(m map[string]map[string]int, path, country string) {
    mm, ok := m[path]
    if !ok {
        mm = make(map[string]int)
        m[path] = mm
    }
    mm[country]++
}
add(hits, "/doc/", "au")
```

On the other hand, a design that uses a single map with a struct key does away with all that complexity:

```Go
type Key struct {
    Path, Country string
}
hits := make(map[Key]int)
```

When an Vietnamese person visits the home page, incrementing (and possibly creating) the appropriate counter is a one-liner:
```Go
hits[Key{"/", "vn"}]++
```

And it's similarly straightforward to see how many Swiss people have read the spec:

```Go
n := hits[Key{"/ref/spec", "ch"}]
```

 Source:  https://blog.golang.org/maps

 ## Big O
 <hr>
  Hash tables in general are O(1) for inserting, looking up and deleting data; we can assume that this is also true for Go's maps.