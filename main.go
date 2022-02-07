package main

import (
	"fmt"
	"lrucache/cache"
)

func main()  {
	withCacheFn, _ := cache.WithInt64Cache(func(key int64) interface{} {
		return key*key
	})

	for i := 0; i < 150; i++ {
		_ = withCacheFn(int64(i))
	}

	fmt.Println("\n\n-------     testing...     --------")
	fmt.Println(withCacheFn(145))
	fmt.Println(withCacheFn(146))
	fmt.Println(withCacheFn(147))
	fmt.Println(withCacheFn(1))
	fmt.Println(withCacheFn(2))
	fmt.Println(withCacheFn(52))
}
