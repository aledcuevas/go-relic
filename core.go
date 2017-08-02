package main

// #include<relic.h>
/*
#define ZERO_C	0
#define ONE_C		1
*/
import "C"
import "errors"

func coreInit() int {
	if C.core_init() != C.STS_OK {
		C.core_clean()
		// Error
		return 1
	}
	// Correct (STS_OK)
	return 0
}

func coreClean() {
	C.core_clean()
}

func paramSet() {
	if C.ep_param_set_any() != C.STS_OK {
		C.core_clean()
		panic(MsgNoCurve)
	}
}

/*
 * gToCflag is used to convert Go ints to C.ints, but only limited to 0 and 1. This is necessary because RELIC uses 1 and 0 as flags for many functions. Furthermore, a general conversion function is desirable but left for future work.
 */
func gToCflag(i int) C.int {
	switch i {
	case 0:
		return C.ZERO_C
	case 1:
		return C.ONE_C
	default:
		panic("bad conversion")
	}
}

type goCore struct {
	pcBytes   int
	goArch    int
	relicArch int
}

func setGoParameters() {

}

func cArch() (int, error) {
	s := C.sizeof_int
	switch s {
	case 2:
		return 2, nil
	case 4:
		return 4, nil
	case 8:
		return 8, nil
	default:
		return 0, errors.New("wrong conversion")
	}
}

func gToC(ar int) (C.int, error) {
	switch ar {
	case 2:
		return 2, nil
	case 4:
		return 4, nil
	case 8:
		return 8, nil
	default:
		return 0, errors.New("architecture not recognized")
	}
}

//Initialize a struct?
//Save parameters in a struct
//func to set RELIC architecture parameters
// 		include values for curves to be used
// 		include prime lengths and stuff
//		alter any existing if defs
