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

func ShortTest(t *testing.T) {

}

func LongTest(t *testing.T) {

}

func TestMemory1(t *testing.T) {
	testBegin("memory can be allocated")
	coreInit()
	epParamSetAny()
	defer coreClean()
	a := new(ellPointG1)
	nullG1(&a.g1)
	newG1(&a.g1)
	freeG1(&a.g1)
}

func TestUtil1(t *testing.T) {
	testBegin("comparison is consistent")
	coreInit()
	epParamSetAny()
	defer coreClean()
	a := new(ellPointG1)
	b := new(ellPointG1)
	c := new(ellPointG1)

	//l, code := StsErr, StsErr
	binSize := 2*PcBytes + 1
	bin := make([]byte, binSize, binSize)

	nullG1(&a.g1)
	nullG1(&b.g1)
	nullG1(&c.g1)

	newG1(&a.g1)
	newG1(&b.g1)
	newG1(&c.g1)

	defer freeG1(&a.g1)
	defer freeG1(&b.g1)
	defer freeG1(&c.g1)

	randG1(&a.g1)
	randG1(&b.g1)

	if epCmp(&a.g1, &b.g1) != CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("copy and comparison are consistent")
	randG1(&c.g1)

	if epCmp(&a.g1, &c.g1) == CmpNe {
		epCopy(&c.g1, &a.g1)
		if epCmp(&c.g1, &a.g1) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}
	}
	if epCmp(&b.g1, &c.g1) == CmpNe {
		epCopy(&c.g1, &b.g1)
		if epCmp(&b.g1, &c.g1) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}
	}

	epDbl(&c.g1, &a.g1)
	epNorm(&c.g1, &c.g1)
	epDbl(&a.g1, &a.g1)

	if epCmp(&c.g1, &a.g1) == CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}
	if epCmp(&a.g1, &c.g1) == CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	epDbl(&c.g1, &c.g1)
	epDbl(&a.g1, &a.g1)
	if epCmp(&c.g1, &a.g1) == CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("inversion and comparison are consistent")

	epRand(&a.g1)
	epNegBasic(&b.g1, &a.g1)
	if epCmp(&a.g1, &b.g1) != CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("assignment to random/infinity and comparison are consistent")

	epRand(&a.g1)
	epSetInfinity(&c.g1)
	if epCmp(&a.g1, &c.g1) != CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}
	if epCmp(&c.g1, &a.g1) != CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("assignment to infinity and infinity test are consistent")

	epSetInfinity(&a.g1)
	if epIsInfinity(&a.g1) {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("reading and writing a point are consistent")
	for i := 0; i < 2; i++ {
		j := int32(i)
		epSetInfinity(&a.g1)
		l := epSizeBin(&a.g1, j)
		epWriteBin(bin, l, &a.g1, j)
		epReadBin(&b.g1, bin, l)

		if epCmp(&a.g1, &b.g1) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}

		epRand(&a.g1)
		l = epSizeBin(&a.g1, j)
		epWriteBin(bin, l, &a.g1, j)
		epReadBin(&b.g1, bin, l)

		if epCmp(&a.g1, &b.g1) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}

		epRand(&a.g1)
		epDbl(&a.g1, &a.g1)
		l = epSizeBin(&a.g1, j)
		epNorm(&a.g1, &a.g1)
		epWriteBin(bin, l, &a.g1, j)
		epReadBin(&b.g1, bin, l)

		if epCmp(&a.g1, &b.g1) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}

	}
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
