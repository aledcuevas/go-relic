package main

// #include <relic.h>
import "C"
import "fmt"

func main() {
	coreInit()
	paramSet()
	var a C.ep_st
	newG1(&a)
	C.ep_rand(&a)

	fmt.Println("aaaaa")
}
