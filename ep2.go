package main

// #include <relic.h>
import "C"

// ******* GENERAL METHODS *******

// // Utils

func ep2Cmp(point1 *C.ep2_st, point2 *C.ep2_st) C.int {
	return C.ep2_cmp(point1, point2)
}

func ep2SetInfinity(point *C.ep2_st) {
	C.ep2_set_infty(point)
}

func ep2IsInfinity(point *C.ep2_st) C.int {
	return C.ep2_is_infty(point)
}

func ep2Rand(point *C.ep2_st) {
	C.ep2_rand(point)
}

func ep22IsValid(point *C.ep2_st) C.int {
	return C.ep2_is_valid(point)
}

func ep2Print(point *C.ep2_st) {
	C.ep2_print(point)
}

func ep2Copy(to *C.ep2_st, from *C.ep2_st) {
	C.ep2_copy(to, from)
}

// // Addition

func ep2AddBasic(result *C.ep2_st, point1 *C.ep2_st, point2 *C.ep2_st) {
	C.ep2_add_basic(result, point1, point2)
}

// // Negation

func ep2NegBasic(result *C.ep2_st, point *C.ep2_st) {
	C.ep2_neg_basic(result, point)
}
