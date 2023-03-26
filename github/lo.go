package github

import (
	"fmt"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func Sli() {
	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
	// []string{"Samuel", "John"}
	fmt.Println(names)

	nums := lo.Uniq[int64]([]int64{6, 666, 1234, 666})
	// []string{6, 666, 1234}
	fmt.Println(nums)

	present := lo.Contains([]int{0, 1, 2, 3, 4, 5}, 5)
	// true
	fmt.Println(present)

	ok := lo.Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	// true
	fmt.Println(ok)

	ok1 := lo.Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	// false
	fmt.Println(ok1)

	lo.ForEach([]string{"hello", "world"}, func(x string, _ int) {
		println(x)
	})
	lop.ForEach([]string{"hello", "world"}, func(x string, _ int) {
		println(x)
	})

	in := []string{"", "foo", "", "bar", ""}

	slice := lo.Compact[string](in)
	// []string{"foo", "bar"}
	fmt.Println(slice)

	mergedMaps := lo.Assign[string, int](
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)
	// map[string]int{"a": 1, "b": 3, "c": 4}
	fmt.Println(mergedMaps)

	reverseOrder := lo.Reverse([]int{0, 1, 2, 3, 4, 5})
	// []int{5, 4, 3, 2, 1, 0}
	fmt.Println(reverseOrder)

}

func Mpa() {
	keys := lo.Keys[string, int](map[string]int{"foo": 1, "bar": 2})
	// []string{"foo", "bar"}
	fmt.Println(keys)

	m := map[int]int64{1: 4, 2: 5, 3: 6}

	s := lo.MapToSlice(m, func(k int, v int64) string {
		return fmt.Sprintf("%d_%d", k, v)
	})
	// []string{"1_4", "2_5", "3_6"}
	fmt.Println(s)

	m1 := lo.Invert(map[string]int{"a": 1, "b": 2})
	// map[int]string{1: "a", 2: "b"}
	fmt.Println(m1)

	m2 := lo.Invert(map[string]int{"a": 1, "b": 2, "c": 1})
	// map[int]string{1: "c", 2: "b"}
	fmt.Println(m2)
}
