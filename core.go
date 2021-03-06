package main

// #include<relic.h>
/*
#define ZERO_C	0
#define ONE_C		1
*/
import "C"

/***
 * core.go contains the set of functions required to instantiate the library.
 * These set of functions should be the first called in a project.
 * RELIC maintains a global ctx_t struct to implement many functions across the board (e.g. TRY/CATCH), this struct is instantiated with core_init.
 * The parameters to be used in a context, are set with the param_set family of functions.
 * In this case, we are using the default ep_param_set_any for parameters in elliptic curves over prime fields.
 * The default security level (i.e. type of curve) is set in the makefile before compilation, but it can also be undef'd and redef'd.
 * For a more comprehensive list of methods and available curves, check out relic_ep_param.c
 * NOTE: ZERO_C and ONE_C are defined to make error checking easier and less error-prone, by doing C to C int comparisons.
 * It would be interesting to define a Go struct with the most important data from ctx_t to facilitate access to info such as size of integers, size of primes, curves, etc.
 ***/

//CoreInit initializes the library. Returns a flag with success/error.
//Found in relic_core.c
func coreInit() int {
	if C.core_init() != C.STS_OK {
		C.core_clean()
		return StsErr
	}
	return StsOk
}

//CoreClean finalizes the library. Returns a flag with success/error.
//Found in relic_core.c
func coreClean() int {
	if C.core_clean() != C.STS_OK {
		return StsErr
	}
	return StsOk
}

//EpParamSetAny configure some set of curve parameters for the current security level.
//Found in relic_ep_param.c
func epParamSetAny() int {
	if C.ep_param_set_any() != C.STS_OK {
		C.core_clean()
		return StsErr
	}
	return StsOk
}
