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
