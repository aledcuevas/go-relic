package main

// #include<relic.h>
import "C"

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
