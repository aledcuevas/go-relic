package main

// #include <relic.h>
/*
void g2_free_w(ep2_st* p){
  g2_free(p);
}
void g2_null_w(ep2_st* p){
  g2_null(p);
}
void g2_new_w(ep2_st* p){
  g2_new(p);
}
void g2_rand_w(ep2_st* p){
  g2_rand(p);
}
int g2_cmp_w(ep2_st* p, ep2_st* q){
  return g2_cmp(p,q);
}
void g2_dbl_w(ep2_st* r, ep2_st* p){
  g2_dbl(r,p);
}
void g2_norm_w(ep2_st* r, ep2_st* p){
  g2_norm(r,p);
}
void g2_neg_w(ep2_st* r, ep2_st* p){
  g2_neg(r, p);
}
void g2_copy_w(ep2_st* t, ep2_st* f){
  g2_copy(t, f);
}
void g2_sub_w(ep2_st* r, ep2_st* p, ep2_st* q){
	g2_sub(r,p,q);
}
void g2_add_w(ep2_st* r, ep2_st* p, ep2_st* q){
  g2_add(r,p,q);
}
void g2_mul_w(ep2_st* r, ep2_st* p, bn_st* k){
  g2_mul(r,p,k);
}
void g2_read_bin_w(ep2_st* r, uint8_t* bin, int len){
	g2_read_bin(r, bin, len);
}
int g2_size_bin_w(ep2_st* p, int pack){
	return g2_size_bin(p, pack);
}
void g2_write_bin_w(uint8_t* rbin, int len, ep2_st* p, int pack){
	g2_write_bin(rbin, len, p, pack);
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

type pointG2 struct {
	g         C.ep2_st
	generator string
}

func newPointG2(gen string) *pointG2 {
	pg := new(pointG2)
	C.g2_new_w(&pg.g)
	pg.generator = gen
	//runtime.SetFinalizer(&pg.g, clear)
	return pg
}

func (p *pointG2) Equal(q kyber.Point) bool {
	pg := q.(*pointG2)
	i := C.g2_cmp_w(&p.g, &pg.g)
	switch i {
	case C.CMP_EQ:
		return true
	case C.CMP_NE:
		return false
	default:
		panic("Error in C casting")
	}
}

func (p *pointG2) Null() kyber.Point {
	C.g2_null_w(&p.g)
	return p
}

func (p *pointG2) Base() kyber.Point {
	//get a base
	panic("not implemented")
}

func (p *pointG2) Add(p1, p2 kyber.Point) kyber.Point {
	pg1 := p1.(*pointG2)
	pg2 := p2.(*pointG2)
	C.g2_add_w(&p.g, &pg1.g, &pg2.g)
	return p
}

func (p *pointG2) Sub(p1, p2 kyber.Point) kyber.Point {
	pg1 := p1.(*pointG2)
	pg2 := p2.(*pointG2)
	C.g2_sub_w(&p.g, &pg1.g, &pg2.g)
	return p
}

func (p *pointG2) Neg(p1 kyber.Point) kyber.Point {
	pg := p1.(*pointG2)
	C.g2_neg_w(&p.g, &pg.g)
	return p
}

func (p *pointG2) Mul(s kyber.Scalar, p1 kyber.Point) kyber.Point {
	//Waiting for a scalar implementation
	return p
}

func (p *pointG2) Clone() kyber.Point {
	p2 := new(pointG2)
	C.g2_new_w(&p2.g)
	C.g2_copy_w(&p2.g, &p.g)
	return p2
}

//******TODO******
func (p *pointG2) Data() ([]byte, error) {
	panic("not implemented")
}

func (p *pointG2) EmbedLen() int {
	panic("not implemented")
}

func (p *pointG2) Embed(data []byte, rand cipher.Stream) kyber.Point {
	panic("not implemented")
}

//***************

func (p *pointG2) Pick(rand cipher.Stream) kyber.Point {
	//TODO: rand is currently not being used because RELIC randomizes a point
	// by reference.
	C.g2_rand_w(&p.g)
	return p
}

func (p *pointG2) PickR() kyber.Point {
	//TODO: rand is currently not being used because RELIC randomizes a point
	// by reference.
	C.g2_rand_w(&p.g)
	return p
}

func (p *pointG2) Set(p2 kyber.Point) kyber.Point {
	pg2 := p2.(*pointG2)
	C.g2_copy_w(&p.g, &pg2.g)
	return p
}

func (p *pointG2) SetVarTime(varTime bool) error {
	return errors.New("ErrVarTime")
}

func (p *pointG2) String() string {
	buff, err := p.MarshalBinary()
	if err != nil {
		panic("Error in marshalling")
	}
	return hex.EncodeToString(buff)
}

// -- -- -- -- Marshalling

func (p *pointG2) MarshalBinary() ([]byte, error) {
	len := p.MarshalSize()
	//Get len from byte size
	//NOTE: I'm making the casting specific so you can have an idea of how to play
	// around with this value.
	clen := C.int(int32(len))
	//Point compression flag
	pack := C.int(0)
	//We obtain a C array allocated in the C heap using buff. CBytes returns an unsafe.Pointer
	buff := make([]byte, len, len)
	b := C.CBytes(buff)
	defer C.free(b)
	//We perform the write operation on the C array, which has to be casted to match the signature
	C.g2_write_bin_w((*C.uint8_t)(b), clen, &p.g, pack)
	//We transform the unsafe.Pointer back to a []byte
	//GoBytes takes C data with explicity length and returns Go []byte
	copy(buff, C.GoBytes(b, clen))
	//TODO: Implement error-checking functionality
	return buff, nil
}

func (p *pointG2) MarshalTo(w io.Writer) (int, error) {
	return pointMarshalTo(p, w)
}

func (p *pointG2) UnmarshalBinary(buff []byte) error {
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
	C.g2_read_bin_w(&p.g, (*C.uint8_t)(b), len)
	//TODO: Implement error-checking functionality
	return nil
}

func (p *pointG2) UnmarshalFrom(r io.Reader) (int, error) {
	return pointUnmarshalFrom(p, r)
}

func (p *pointG2) MarshalSize() int {
	//Point compression flag
	pack := C.int(0)
	return int(C.g2_size_bin_w(&p.g, pack))
}
