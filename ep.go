package main

// #include <relic.h>
// #include <stdint.h>
/*
uint8_t* btou8(char* b, int size){
	uint8_t* r = malloc(size);
	memcpy(r, b, size);
	return &r[0];
}
*/
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

/*** CODE BELOW NEEDS STRONG REVIEW ***/
//TODO: Need to make sure memory writing errors are caught
/*
 * epSizeBin returns the size of a point with or without compression (pack). It is not desirable to expose C types in Go function returns. However, due to the lack of a conversion function, it is advised to take the result and "plug" it in subsequent function calls as necessary.
 */
func epSizeBin(point *C.ep_st, pack int32) int32 {
	p := C.int(pack)
	size := C.ep_size_bin(point, p)
	return int32(size)
}

func epReadBin(rPoint *C.ep_st, bin []byte, len int32) {
	/* Go []byte slice to C array -- returns an unsafe.Pointer */
	b := C.CBytes(bin)
	defer C.free(b)
	/* We need to match the uint8_t signature in ep_read_bin. we cast tmp to
	a char* and copy the bytes directly to a malloc'd block. we return the address
	of this block of memory with the right signature. Note: the casting of
	uint to C.int has not yet been thoroughly tested.
	*/
	//b := C.btou8((*C.char)(unsafe.Pointer(&tmp)), (C.int)(unsafe.Sizeof(tmp)))
	//defer C.free(unsafe.Pointer(&b))
	l := C.int(len)

	C.ep_read_bin(rPoint, (*C.uint8_t)(b), l)
}

func epWriteBin(rBin []byte, len int32, point *C.ep_st, pack int32) {
	p := C.int(pack)
	l := C.int(len)
	//Type conversion from int32 to size_t not thoroughly tested
	lun := C.size_t(len)
	addr := C.malloc(lun)
	defer C.free(addr)

	C.ep_write_bin((*C.uint8_t)(addr), l, point, p)

	//TODO fix this assignment
	rBin = C.GoBytes(addr, l)

}

/*** CODE ABOVE NEEDS STRONG REVIEW ***/
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
