package main

import (
	"fmt"
	"github.com/golang/groupcache/lru"
)

func main() {
	fmt.Println("测试lru")
	cache :=lru.New(20);

	for i:=0;i<18;i++{
		cache.Add(i,i)
	}

	fmt.Println(cache.Len())
	cache.Add(1,2222)
	cache.Add(1,21222)
	cache.Add(1,22222)
	fmt.Println(cache.Len())

}
