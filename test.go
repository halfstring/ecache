package main

import (
	ecache "./src"
	"time"
	"fmt"
	"log"
)

type myData struct {
	value string
}

func main() {
	cache := ecache.Cache("default")

	cache.Add("kalep", &myData{value: "kalepvalue"}, 50*time.Second)

	v,ok:= cache.Value("kalep")
	if ok != nil {
		log.Println(ok)
	}

	fmt.Println(v)
}
