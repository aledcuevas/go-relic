package main

import (
	"fmt"
	"testing"
)

/***
 * List of tests given by RELIC but re-implemented using the Go methods.
 * Test files in Go can't utilize Cgo.
 * The following tests were taken from test_pc.c, a test suite for pairing crypto.
 * NOTE: The if/else statements are ugly but did the job.
 ***/

func TestMemory1(t *testing.T) {
}

func TestUtil1(t *testing.T) {
}

// *** SUPPLEMENTARY FUNCTIONS ***

func testBegin(s string) {
	fmt.Printf("Testing if %v\n", s)
}

func passed() {
	fmt.Println("-PASS-")
}
func failed() {
	fmt.Println("-FAIL-")
}
