# collections

[![Go Reference](https://pkg.go.dev/badge/github.com/ariolwork/collections.svg)](https://pkg.go.dev/github.com/ariolwork/collections)

Go package `collections` provides helpers to exec operation with all built-in collection types: slices, maps and iters.

```
go get github.com/ariolwork/collections
```

API is stable. [Documentation](https://pkg.go.dev/github.com/ariolwork/collections).

```go
package example

import (
	"encoding/json"

	"github.com/ariolwork/collections"
)

type Address struct {
	Index int
    Name    *string
	ChildAddress *string
}

// build map from deserialized slice collection
var addressesMap = slices.ToMap(addresses, func(a Address) int {return a.Index})

func Search(index int) (Address, error) () {
    if add, ok := addressesMap[index]; ok {
        return add, nil
    }

    return Address{}, fmt.Errorf("Unknown index: %d", index)
}
```