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
