package main

import "testing"

func ShortTest(t *testing.T) {

}

func LongTest(t *testing.T) {

}

func TestMemory1(t *testing.T) {
	testBegin("memory can be allocated")
	a := new(ellPointG1)
	nullG1(&a.g1)
	newG1(&a.g1)
	freeG1(&a.g1)
}

func TestUtil1(t *testing.T) {
	testBegin("comparison is consistent")
	a := new(ellPointG1)
	b := new(ellPointG1)
	c := new(ellPointG1)

	nullG1(&a.g1)
	nullG1(&b.g1)
	nullG1(&c.g1)

}
