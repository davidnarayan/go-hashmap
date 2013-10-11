go-hashmap
==========

go-hashmap provides a mechanism to log maps as a string of key/value pairs

```go
package main

import (
	"fmt"
	"github.com/davidnarayan/go-hashmap"
)

func main() {
	m := make(map[string]interface{})
	hm := hashmap.NewHashMap(m)
	hm.Set("foo", "123")
	hm.Set("bar", "456")

	fmt.Println(hm)
}
```

This outputs:

```
bar="456" foo="123"
```
