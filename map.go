package main

import (
	"fmt"
	)
	
func main() {

var m = make(map[string]int)
m["a"] =1
fmt.Println(m)
//map[a:1]

fmt.Println(m["b"])
showFromMap(m)
//0
}



func showFromMap (m map[string]int){
    fmt.Println(m["b"])
}