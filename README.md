go-hashmap
==========

go-hashmap provides a way to log maps as a string of key/value pairs

```go
package main

import (
	"fmt"
	"github.com/davidnarayan/go-hashmap"
)

func main() {
	hm := hashmap.NewHashMap()
	hm.Set("foo", "123")
	hm.Set("bar", "456")

	fmt.Println(hm)
}
```

This outputs:

```
bar="456" foo="123"
```
