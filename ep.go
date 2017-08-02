package main

// #include <relic.h>
import "C"

// ******* GENERAL METHODS *******

// // Utils

func epCmp(point1 *C.ep_st, point2 *C.ep_st) int {
	var result int
	switch C.ep_cmp(point1, point2) {
	case C.CMP_NE:
		result = CmpNe
	case C.CMP_EQ:
		result = CmpEq
	}
	return result
}

func epSetInfinity(point *C.ep_st) {
	C.ep_set_infty(point)
}

func epIsInfinity(point *C.ep_st) bool {
	var result bool
	/* CMP_GT and CMP_EQ are used because they are constants in the RELIC library that represent '1' and '0', respectively.*/
	switch C.ep_is_infty(point) {
	case C.CMP_GT:
		result = true
	case C.CMP_EQ:
		result = false
	}
	return result
}

func epRand(point *C.ep_st) {
	C.ep_rand(point)
}

func epIsValid(point *C.ep_st) bool {
	var result bool
	switch C.ep_is_valid(point) {
	case C.CMP_GT:
		result = true
	case C.CMP_EQ:
		result = false
	}
	return result
}

func epPrint(point *C.ep_st) {
	C.ep_print(point)
}

func epCopy(to *C.ep_st, from *C.ep_st) {
	C.ep_copy(to, from)
}

/*
 * epSizeBin returns the size of a point with or without compression (pack). It is not desirable to expose C types in Go function returns. However, due to the lack of a conversion function, it is advised to take the result and "plug" it in subsequent function calls as necessary.
 */
func epSizeBin(point *C.ep_st, pack int) C.int {
	p := gToCflag(pack)
	size := C.ep_size_bin(point, p)
	return size
}

func epReadBin(point *C.ep_st, bin int, len int) {

}

func epWriteBin(bin *int, len int, point *C.ep_st, pack int) {

}

// // Addition

func epAddBasic(result *C.ep_st, point1 *C.ep_st, point2 *C.ep_st) {
	C.ep_add_basic(result, point1, point2)
}

// // Negation

/*g1_neg -> ep_neg -> ep_neg_basic*/
func epNegBasic(result *C.ep_st, point *C.ep_st) {
	C.ep_neg_basic(result, point)
}

// // Other Operations

func epDbl(result *C.ep_st, point *C.ep_st) {
	C.ep_dbl_basic(result, point)
}

func epNorm(result *C.ep_st, point *C.ep_st) {
	C.ep_norm(result, point)
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

func setGoParameters() {

}

//Initialize a struct?
//Save parameters in a struct
//func to set RELIC architecture parameters
// 		include values for curves to be used
// 		include prime lengths and stuff
//		alter any existing if defs
