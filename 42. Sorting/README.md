
# Sorting

### Source

```
package main

import (
    "fmt"
    "sort"
)

func main() {

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
```
### Output
```
$ go run sorting.go
Strings: [a b c]
Ints:    [2 4 7]
Sorted:  true
```


Go’s sort package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first.


Sort methods are specific to the builtin type; here’s an example for strings. Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
	


An example of sorting ints.
	


We can also use sort to check if a slice is already in sorted order.
	


Running our program prints the sorted string and int slices and true as the result of our AreSorted test.
