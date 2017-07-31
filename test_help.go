package main

// #include <relic.h>
import "C"
import "fmt"

func testBegin(s string) {
	fmt.Printf("Testing if %v\n", s)
	coreInit()
	paramSet()
}

func passed() {
	fmt.Println("-PASS-")
}
func failed() {
	fmt.Println("-FAIL-")
}
