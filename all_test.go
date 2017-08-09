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
 * BUG: ep2 tests fail, probably because of ep2_rand()
 ***/

func TestMemory1(t *testing.T) {
	fmt.Println("********** G1 TESTS BEGIN **********")
	testBegin("core can be initialized")
	if coreInit() == StsErr {
		t.FailNow()
	}
	if epParamSetAny() == StsErr {
		t.FailNow()
	}
	defer coreClean()
	epParamPrint()
	testBegin("memory can be allocated")
	a := new(pointG1)
	nullG1(&a.g1)
	newG1(&a.g1)
	freeG1(&a.g1)
}

func TestUtil1(t *testing.T) {
	testBegin("comparison is consistent")
	coreInit()
	epParamSetAny()
	defer coreClean()
	a := new(pointG1)
	b := new(pointG1)
	c := new(pointG1)

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

//****************************** //
//********** G2 TESTS ********** //
//****************************** //
/*
func TestMemory2(t *testing.T) {
	fmt.Println("********** G2 TESTS BEGIN **********")
	coreInit()
	epParamSetAny()
	defer coreClean()
	testBegin("memory can be allocated")
	a := new(pointG2)
	nullG2(&a.g2)
	newG2(&a.g2)
	freeG2(&a.g2)
}

func TestUtil2(t *testing.T) {
	testBegin("comparison is consistent")
	coreInit()
	epParamSetAny()
	defer coreClean()
	a := new(pointG2)
	b := new(pointG2)
	c := new(pointG2)

	//l, code := StsErr, StsErr
	binSize := 4*PcBytes + 1
	bin := make([]byte, binSize, binSize)

	nullG2(&a.g2)
	nullG2(&b.g2)
	nullG2(&c.g2)

	newG2(&a.g2)
	newG2(&b.g2)
	newG2(&c.g2)

	defer freeG2(&a.g2)
	defer freeG2(&b.g2)
	defer freeG2(&c.g2)

	randG2(&a.g2)
	randG2(&b.g2)

	if ep2Cmp(&a.g2, &b.g2) != CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("copy and comparison are consistent")
	randG2(&c.g2)

	if ep2Cmp(&a.g2, &c.g2) == CmpNe {
		ep2Copy(&c.g2, &a.g2)
		if ep2Cmp(&c.g2, &a.g2) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}
	}
	if ep2Cmp(&b.g2, &c.g2) == CmpNe {
		ep2Copy(&c.g2, &b.g2)
		if ep2Cmp(&b.g2, &c.g2) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}
	}

	ep2Dbl(&c.g2, &a.g2)
	ep2Norm(&c.g2, &c.g2)
	ep2Dbl(&a.g2, &a.g2)

	if ep2Cmp(&c.g2, &a.g2) == CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}
	if ep2Cmp(&a.g2, &c.g2) == CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	ep2Dbl(&c.g2, &c.g2)
	ep2Dbl(&a.g2, &a.g2)
	if ep2Cmp(&c.g2, &a.g2) == CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("inversion and comparison are consistent")

	ep2Rand(&a.g2)
	ep2NegBasic(&b.g2, &a.g2)
	if ep2Cmp(&a.g2, &b.g2) != CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("assignment to random/infinity and comparison are consistent")

	ep2Rand(&a.g2)
	ep2SetInfinity(&c.g2)
	if ep2Cmp(&a.g2, &c.g2) != CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}
	if ep2Cmp(&c.g2, &a.g2) != CmpEq {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("assignment to infinity and infinity test are consistent")

	ep2SetInfinity(&a.g2)
	if ep2IsInfinity(&a.g2) {
		passed()
	} else {
		failed()
		t.Fail()
	}

	testBegin("reading and writing a point are consistent")
	for i := 0; i < 2; i++ {
		j := int32(i)
		ep2SetInfinity(&a.g2)
		l := ep2SizeBin(&a.g2, j)
		ep2WriteBin(bin, l, &a.g2, j)
		ep2ReadBin(&b.g2, bin, l)

		if ep2Cmp(&a.g2, &b.g2) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}

		ep2Rand(&a.g2)
		l = ep2SizeBin(&a.g2, j)
		ep2WriteBin(bin, l, &a.g2, j)
		ep2ReadBin(&b.g2, bin, l)

		if ep2Cmp(&a.g2, &b.g2) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}

		ep2Rand(&a.g2)
		ep2Dbl(&a.g2, &a.g2)
		l = ep2SizeBin(&a.g2, j)
		ep2Norm(&a.g2, &a.g2)
		ep2WriteBin(bin, l, &a.g2, j)
		ep2ReadBin(&b.g2, bin, l)

		if ep2Cmp(&a.g2, &b.g2) == CmpEq {
			passed()
		} else {
			failed()
			t.Fail()
		}

	}
}
*/
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
