package main

import (
	"time"
	"log"
	"fmt"
	ecache "./src"
)

type myData struct {
	value string
}

func main() {

	//假如 想根据map的key来动态实例化的结构体,然后修改结构体的值，比如:    //obj := m[1]    //obj.One = 10
	cache := ecache.Cache("default")

	cache.Add("kalep", &myData{value: "kalepvalue"}, 50*time.Second)

	v, ok := cache.Value("kalep")
	if ok != nil {
		log.Println(ok)
	}

	fmt.Println(v)
}
