package main

// #include<relic.h>
import "C"

func coreInit() int {
	if C.core_init() != C.STS_OK {
		C.core_clean()
		return 1 // Error
	}
	return 0 // Correct (STS_OK)
}

func coreClean() {
	C.core_clean()
}

func paramSet() {
	if C.ep_param_set_any() != C.STS_OK {
		C.core_clean()
		panic("no curve supported at this security level")
	}
}
