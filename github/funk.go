package github

import "github.com/thoas/go-funk"

type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f Foo) TableName() string {
	return "foo"
}

func Contains() {
	f := &Foo{
		ID:        1,
		FirstName: "Foo",
		LastName:  "Bar",
		Age:       30,
	}

	funk.Contains([]string{"foo", "bar"}, "bar") // true
	// slice of Foo ptr
	funk.Contains([]*Foo{f}, f) // true
	funk.Contains([]*Foo{f}, func(foo *Foo) bool {
		return foo.ID == f.ID
	}) // true
	funk.Contains([]*Foo{f}, nil) // false

	b := &Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       28,
	}

	funk.Contains([]*Foo{f}, b) // false

	// string
	funk.Contains("florent", "rent") // true
	funk.Contains("florent", "foo")  // false

	// even map
	funk.Contains(map[int]string{1: "Florent"}, 1) // true
	funk.Contains(map[int]string{1: "Florent"}, func(key int, name string) bool {
		return key == 1 // or `name == "Florent"` for the value type
	}) // true
}
