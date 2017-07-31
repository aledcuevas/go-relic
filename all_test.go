package main

import "testing"

func ShortTest(t *testing.T) {

}

func LongTest(t *testing.T) {

}

func TestMemory1(t *testing.T) {
	testBegin("memory can be allocated")
	defer coreClean()
	a := new(ellPointG1)
	nullG1(&a.g1)
	newG1(&a.g1)
	freeG1(&a.g1)
}

func TestUtil1(t *testing.T) {
	testBegin("comparison is consistent")
	defer coreClean()
	a := new(ellPointG1)
	b := new(ellPointG1)
	c := new(ellPointG1)

	nullG1(&a.g1)
	nullG1(&b.g1)
	nullG1(&c.g1)

	newG1(&a.g1)
	newG1(&b.g1)
	newG1(&c.g1)

	randG1(&a.g1)
	randG1(&b.g1)
	randG1(&c.g1)
	if epCmp(&a.g1, &c.g1) == CmpNe {
		epCopy(&c.g1, &a.g1)
		if epCmp(&c.g1, &a.g1) == CmpEq {
			passed()
		} else {
			failed()
		}
	}
	if epCmp(&b.g1, &c.g1) == CmpNe {
		epCopy(&c.g1, &b.g1)
		if epCmp(&b.g1, &c.g1) == CmpEq {
			passed()
		} else {
			failed()
		}
	}

}
