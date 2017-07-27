package main

// #include <relic.h>
import "C"

// ******* GENERAL METHODS *******

// // Utils

func epCmp(point1 *C.ep_st, point2 *C.ep_st) C.int {
	return C.ep_cmp(point1, point2)
}

func epSetInfinity(point *C.ep_st) {
	C.ep_set_infty(point)
}

func epIsInfinity(point *C.ep_st) C.int {
	return C.ep_is_infty(point)
}

func epRand(point *C.ep_st) {
	C.ep_rand(point)
}

func epIsValid(point *C.ep_st) C.int {
	return C.ep_is_valid(point)
}

func epPrint(point *C.ep_st) {
	C.ep_print(point)
}

func epCopy(to *C.ep_st, from *C.ep_st) {
	C.ep_copy(to, from)
}

// // Addition

func epAddBasic(result *C.ep_st, point1 *C.ep_st, point2 *C.ep_st) {
	C.ep_add_basic(result, point1, point2)
}

// // Negation

func epNegBasic(result *C.ep_st, point *C.ep_st) {
	C.ep_neg_basic(result, point)
}
