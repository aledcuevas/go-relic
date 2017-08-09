package main

// #include <relic.h>
// #include <stdint.h>
// #include <stdlib.h>
/*
#define ZERO_C	0
#define ONE_C		1
*/
import "C"

// // Utils - relic_ep2_util.c

//Ep2Cmp compares two prime elliptic curve points. Returns eq/neq.
func ep2Cmp(point1 *C.ep2_st, point2 *C.ep2_st) int {
	var result int
	switch C.ep2_cmp(point1, point2) {
	case C.CMP_NE:
		result = CmpNe
	case C.CMP_EQ:
		result = CmpEq
	}
	return result
}

//EpSetInfinity assigns a prime elliptic curve point at the infinity.
func ep2SetInfinity(point *C.ep2_st) {
	C.ep2_set_infty(point)
}

//EpIsInfinity tests if a point on a prime elliptic curve is at the infinity. Returns true/false.
func ep2IsInfinity(point *C.ep2_st) bool {
	var result bool
	switch C.ep2_is_infty(point) {
	case C.ONE_C:
		result = true
	case C.ZERO_C:
		result = false
	}
	return result
}

//EpRand assigns a random value to a prime elliptic curve point.
func ep2Rand(point *C.ep2_st) {
	C.ep2_rand(point)
}

//EpIsValid tests if a point is in the curve. Returns true/false.
func ep2IsValid(point *C.ep2_st) bool {
	var result bool
	switch C.ep2_is_valid(point) {
	case C.CMP_GT:
		result = true
	case C.CMP_EQ:
		result = false
	}
	return result
}

//EpPrint prints a prime elliptic curve point.
func ep2Print(point *C.ep2_st) {
	C.ep2_print(point)
}

//EpCopy copies the second argument to the first argument.
func ep2Copy(to *C.ep2_st, from *C.ep2_st) {
	C.ep2_copy(to, from)
}

/*** CODE BELOW NEEDS STRONG REVIEW ***/
//TODO: Need to make sure memory writing errors are caught
/*
 * Ep2SizeBin returns the size of a point with or without compression (pack). It is not desirable to expose C types in Go function returns. However, due to the lack of a conversion function, it is advised to take the result and "plug" it in subsequent function calls as necessary.
 */
//EpSizeBin returns the number of bytes necessary to store a prime elliptic curve point with optional point compression.
func ep2SizeBin(point *C.ep2_st, pack int32) int32 {
	p := C.int(pack)
	size := C.ep2_size_bin(point, p)
	return int32(size)
}

//EpReadBin reads a prime elliptic curve point from a byte vector in big-endian format.
func ep2ReadBin(rPoint *C.ep2_st, bin []byte, len int32) {
	l := C.int(len)
	//Go []byte slice to C array -- returns an unsafe.Pointer
	//RELIC uses a uint8 array to rep2resent a byte array. Therefore, we will take
	//a byte slice, convert it to a C array, cast it to have the right signature,
	//and pass a pointer to it.
	b := C.CBytes(bin)
	defer C.free(b)
	C.ep2_read_bin(rPoint, (*C.uint8_t)(b), l)
}

//EpWriteBin writes a prime elliptic curve point to a byte vector in bid-endian format with optional point compression.
func ep2WriteBin(rBin []byte, len int32, point *C.ep2_st, pack int32) {
	p := C.int(pack)
	l := C.int(len)
	//We obtain a C array allocated in the C heap using rBin. CBytes returns an unsafe.Pointer
	b := C.CBytes(rBin)
	defer C.free(b)
	//We perform the write operation on the C array, which has to be casted to match the signature
	C.ep2_write_bin((*C.uint8_t)(b), l, point, p)
	//We transform the unsafe.Pointer back to a []byte
	//GoBytes takes C data with explicity length and returns Go []byte
	copy(rBin, C.GoBytes(b, l))

}

/*** CODE ABOVE NEEDS STRONG REVIEW ***/

// // Addition - relic_ep2_add.c

func ep2AddBasic(result *C.ep2_st, point1 *C.ep2_st, point2 *C.ep2_st) {
	C.ep2_add_basic(result, point1, point2)
}

// // Negation - relic_ep2_neg.c

//EpNegBasic negates a prime elliptic curve point represented by affine coordinates.
func ep2NegBasic(result *C.ep2_st, point *C.ep2_st) {
	/*gX_neg -> epX_neg -> epX_neg_basic*/
	C.ep2_neg_basic(result, point)
}

// // Doubling - relic_ep2_dbl.c

//EpDbl doubles a prime elliptic curve point represented in affine coordinates.
func ep2Dbl(result *C.ep2_st, point *C.ep2_st) {
	C.ep2_dbl_basic(result, point)
}

// // Normalization - relic_ep2_norm.c

//EpNorm converts a point to affine coordinates.
func ep2Norm(result *C.ep2_st, point *C.ep2_st) {
	C.ep2_norm(result, point)
}
