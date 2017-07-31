package main

import "testing"

func ShortTest(t *testing.T) {

}

func LongTest(t *testing.T) {

}

func TestMemory1(t *testing.T) {
	testBegin("memory can be allocated")
	coreInit()
	paramSet()
	defer coreClean()
	a := new(ellPointG1)
	nullG1(&a.g1)
	newG1(&a.g1)
	freeG1(&a.g1)
}

func TestUtil1(t *testing.T) {
	testBegin("comparison is consistent")
	coreInit()
	paramSet()
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

	if epCmp(&a.g1, &b.g1) != CmpEq {
		passed()
	} else {
		failed()
	}

	testBegin("copy and comparison are consistent")
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

	epDbl(&c.g1, &a.g1)
	epNorm(&c.g1, &c.g1)
	epDbl(&a.g1, &a.g1)

	if epCmp(&c.g1, &a.g1) == CmpEq {
		passed()
	} else {
		failed()
	}
	if epCmp(&a.g1, &c.g1) == CmpEq {
		passed()
	} else {
		failed()
	}

	epDbl(&c.g1, &c.g1)
	epDbl(&a.g1, &a.g1)
	if epCmp(&c.g1, &a.g1) == CmpEq {
		passed()
	} else {
		failed()
	}

	testBegin("inversion and comparison are consistent")

	epRand(&a.g1)
	epNegBasic(&b.g1, &a.g1)
	if epCmp(&a.g1, &b.g1) != CmpEq {
		passed()
	} else {
		failed()
	}

	testBegin("assignment to random/infinity and comparison are consistent")

	epRand(&a.g1)
	epSetInfinity(&c.g1)
	if epCmp(&a.g1, &c.g1) != CmpEq {
		passed()
	} else {
		failed()
	}
	if epCmp(&c.g1, &a.g1) != CmpEq {
		passed()
	} else {
		failed()
	}

	testBegin("assignment to infinity and infinity test are consistent")

	epSetInfinity(&a.g1)
	if epIsInfinity(&a.g1) {
		passed()
	} else {
		failed()
	}

	testBegin("reading and writing a point are consistent")
	for j := 0; j < 2; j++ {
		epSetInfinity(&a.g1)
	}
}
