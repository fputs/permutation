# permutation
A simple permutation package using generics. Requires go1.18 or higher

## Install
```bash
go get github.com/fputs/permutation@latest
```

## Usage

```go
package main

import (
	"fmt"

	perm "github.com/fputs/permutation"
)

func main() {
	a := []int{1, 2, 3}
	p := perm.Permutations(a)
	fmt.Println(p)
}
```
result:
```
[1 2 3]
[2 1 3]
[3 1 2]
[1 3 2]
[2 3 1]
[3 2 1]
```
