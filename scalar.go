package main

//#include<relic.h>
/*
void bn_free_w(bn_st* s){
  bn_free(s);
}
void bn_null_w(bn_st* s){
  bn_null(s);
}
void bn_new_w(bn_st* s){
  bn_new(s);
}
void bn_rand_w(bn_st* s, int sign, int bits){
  bn_rand(s,sign,bits);
}
void bn_zero_w(bn_st* s){
  bn_zero(s);
}
void bn_set_dig_w(bn_st* s, int d){
  bn_set_dig(s, d);
}
int bn_cmp_w(bn_st* s, bn_st* t){
  return bn_cmp(s,t);
}
void bn_neg_w(bn_st* s, bn_st* t){
  bn_neg(s,t);
}
void bn_copy_w(bn_st* s, bn_st* t){
  bn_copy(s,t);
}
void bn_sub_w(bn_st* r, bn_st* p, bn_st* q){
	bn_sub(r,p,q);
}
void bn_add_w(bn_st* r, bn_st* p, bn_st* q){
  bn_add(r,p,q);
}
void bn_mul_w(bn_st* r, bn_st* p, bn_st* k){
  bn_mul(r,p,k);
}
void bn_div_w(bn_st* r, bn_st* p, bn_st* q){
  bn_add(r,p,q);
}
void bn_read_bin_w(bn_st* r, uint8_t* bin, int len){
	bn_read_bin(r, bin, len);
}
int bn_size_bin_w(bn_st* p){
	return bn_size_bin(p);
}
void bn_write_bin_w(uint8_t* rbin, int len, bn_st* p){
	bn_write_bin(rbin, len, p);
}
*/
import "C"

import (
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"io"

	"gopkg.in/dedis/kyber.v1"
)

type scalar struct {
	fe C.bn_st
}

// NewScalar returns a non initialized kyber.Scalar implementation to use with
// PBC groups.
func NewScalar() kyber.Scalar {
	s := new(scalar)
	C.bn_new_w(&s.fe)
	//runtime.SetFinalizer(s, clearScalar)
	return s
}

func (s *scalar) Zero() kyber.Scalar {
	C.bn_zero_w(&s.fe)
	return s
}

func (s *scalar) One() kyber.Scalar {
	C.bn_null_w(&s.fe)
	return s
}

func (s *scalar) Equal(s2 kyber.Scalar) bool {
	sg := s2.(*scalar)
	i := C.bn_cmp_w(&s.fe, &sg.fe)
	switch i {
	case C.CMP_EQ:
		return true
	case C.CMP_NE:
		return false
	default:
		panic("Error in C casting")
	}
}

func (s *scalar) Neg(s2 kyber.Scalar) kyber.Scalar {
	C.bn_neg(&s.fe, &s2.(*scalar).fe)
	return s
}

func (s *scalar) Add(s1, s2 kyber.Scalar) kyber.Scalar {
	sc1 := s1.(*scalar)
	sc2 := s2.(*scalar)
	C.bn_add_w(&s.fe, &sc1.fe, &sc2.fe)
	return s
}

func (s *scalar) Sub(s1, s2 kyber.Scalar) kyber.Scalar {
	sc1 := s1.(*scalar)
	sc2 := s2.(*scalar)
	C.bn_sub_w(&s.fe, &sc1.fe, &sc2.fe)
	return s
}

func (s *scalar) Mul(s1, s2 kyber.Scalar) kyber.Scalar {
	sc1 := s1.(*scalar)
	sc2 := s2.(*scalar)
	C.bn_mul_w(&s.fe, &sc1.fe, &sc2.fe)
	return s
}

func (s *scalar) Div(s1, s2 kyber.Scalar) kyber.Scalar {
	sc1 := s1.(*scalar)
	sc2 := s2.(*scalar)
	C.bn_div_w(&s.fe, &sc1.fe, &sc2.fe)
	return s
}

func (s *scalar) Inv(s2 kyber.Scalar) kyber.Scalar {
	sc2 := s2.(*scalar)
	//TODO:
	return s
}

func (s *scalar) SetInt64(i int64) kyber.Scalar {
	//The constant must fit on a multiple precision digit, or dig_t type
	// using only the number of bits specified on BN_DIGIT.
	//BN_DIGIT -> DIGIT -> WORD
	//NOTE: WORD is a constant which is set depending on the processor's word size at
	// compile time. It is important to check whether the WORD size is 64 to
	// avoid errors.
	C.bn_set_dig_w(&s.fe, C.int(i))
	return s

}

func (s *scalar) Set(a kyber.Scalar) kyber.Scalar {
	C.bn_copy_w(&s.fe, &a.(*scalar).fe)
	return s
}

func (s *scalar) Clone() kyber.Scalar {
	s2 := NewScalar()
	C.bn_copy_w(&s2.(*scalar).fe, &s.fe)
	return s2
}

func (s *scalar) MarshalBinary() ([]byte, error) {
	len := s.MarshalSize()
	//Get len from byte size
	//NOTE: I'm making the casting specific so you can have an idea of how to play
	// around with this value.
	clen := C.int(int32(len))
	//We obtain a C array allocated in the C heap using buff. CBytes returns an unsafe.Pointer
	buff := make([]byte, len, len)
	b := C.CBytes(buff)
	defer C.free(b)
	//We perform the write operation on the C array, which has to be casted to match the signature
	C.bn_write_bin_w((*C.uint8_t)(b), clen, &s.fe)
	//We transform the unsafe.Pointer back to a []byte
	//GoBytes takes C data with explicity length and returns Go []byte
	copy(buff, C.GoBytes(b, clen))
	//TODO: Implement error-checking functionality
	return buff, nil
}

func (s *scalar) MarshalTo(w io.Writer) (int, error) {
	return scalarMarshalTo(s, w)
}

func (s *scalar) UnmarshalBinary(buff []byte) error {
	//Get len from byte size
	//NOTE: I'm making the casting specific so you can have an idea of how to play
	// around with this value.
	len := C.int(int32(len(buff)))
	//Go []byte slice to C array -- returns an unsafe.Pointer
	//RELIC uses a uint8 array to represent a byte array. Therefore, we will take
	// a byte slice, convert it to a C array, cast it to have the right signature,
	// and pass a pointer to it.
	b := C.CBytes(buff)
	defer C.free(b)
	C.bn_read_bin_w(&s.fe, (*C.uint8_t)(b), len)
	//TODO: Implement error-checking functionality
	return nil
}

func (s *scalar) UnmarshalFrom(r io.Reader) (int, error) {
	return scalarUnmarshalFrom(s, r)
}

func (s *scalar) MarshalSize() int {
	i := int(C.bn_size_bin_w(&s.fe))
	return i
}

func (s *scalar) SetBytes(buff []byte) kyber.Scalar {
	// XXX Maybe later have a "real" setbytes
	s.UnmarshalBinary(buff)
	return s
}

func (s *scalar) Bytes() []byte {
	buff, _ := s.MarshalBinary()
	return buff
}

func (s *scalar) Pick(rand cipher.Stream) kyber.Scalar {
	//TODO: rand is currently not being used because RELIC randomizes a point
	// by reference.
	C.bn_rand_w(&s.fe, C.BN_POS, C.RELIC_BN_BITS)
	return s
}

func (s *scalar) String() string {
	buff, err := s.MarshalBinary()
	if err != nil {
		panic("Error in marshalling")
	}
	return hex.EncodeToString(buff)
}

// clearScalar frees the memory allocated by the C library.
func clearScalar(s *scalar) {
	C.bn_free_w(&s.fe)
}

func (s *scalar) SetVarTime(varTime bool) error {
	return errors.New("ErrVarTime")
}

// -- -- Helper Functions

// ScalarMarshalTo provides a generic implementation of Scalar.EncodeTo
// based on Scalar.Encode.
func scalarMarshalTo(s kyber.Scalar, w io.Writer) (int, error) {
	buf, err := s.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return w.Write(buf)
}

// ScalarUnmarshalFrom provides a generic implementation of Scalar.DecodeFrom,
// based on Scalar.Decode, or Scalar.Pick if r is a Cipher or cipher.Stream.
// The returned byte-count is valid only when decoding from a normal Reader,
// not when picking from a pseudorandom source.
func scalarUnmarshalFrom(s kyber.Scalar, r io.Reader) (int, error) {
	if strm, ok := r.(cipher.Stream); ok {
		s.Pick(strm)
		return -1, nil // no byte-count when picking randomly
	}
	buf := make([]byte, s.MarshalSize())
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return n, err
	}
	return n, s.UnmarshalBinary(buf)
}
