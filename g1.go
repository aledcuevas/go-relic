package main

// #include <relic.h>
/*
void g1_free_w(ep_st* p){
  g1_free(p);
}
void g1_null_w(ep_st* p){
  g1_null(p);
}
void g1_new_w(ep_st* p){
  g1_new(p);
}
void g1_rand_w(ep_st* p){
  g1_rand(p);
}
int g1_cmp_w(ep_st* p, ep_st* q){
  return g1_cmp(p,q);
}
void g1_dbl_w(ep_st* r, ep_st* p){
  g1_dbl(r,p);
}
void g1_norm_w(ep_st* r, ep_st* p){
  g1_norm(r,p);
}
void g1_neg_w(ep_st* r, ep_st* p){
  g1_neg(r, p);
}
void g1_copy_w(ep_st* t, ep_st* f){
  g1_copy(t, f);
}
void g1_sub_w(ep_st* r, ep_st* p, ep_st* q){
	g1_sub(r,p,q);
}
void g1_add_w(ep_st* r, ep_st* p, ep_st* q){
  g1_add(r,p,q);
}
void g1_mul_w(ep_st* r, ep_st* p, bn_st* k){
  g1_mul(r,p,k);
}
void g1_read_bin_w(ep_st* r, uint8_t* bin, int len){
	g1_read_bin(r, bin, len);
}
int g1_size_bin_w(ep_st* p, int pack){
	return g1_size_bin(p, pack);
}
void g1_write_bin_w(uint8_t* rbin, int len, ep_st* p, int pack){
	g1_write_bin(rbin, len, p, pack);
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

type pointG1 struct {
	g         C.ep_st
	generator string
}

func newPointG1(gen string) *pointG1 {
	pg := new(pointG1)
	C.g1_new_w(&pg.g)
	pg.generator = gen
	//runtime.SetFinalizer(&pg.g, clear)
	return pg
}

func (p *pointG1) Equal(q kyber.Point) bool {
	pg := q.(*pointG1)
	i := C.g1_cmp_w(&p.g, &pg.g)
	switch i {
	case C.CMP_EQ:
		return true
	case C.CMP_NE:
		return false
	default:
		panic("Error in C casting")
	}
}

func (p *pointG1) Null() kyber.Point {
	C.g1_null_w(&p.g)
	return p
}

func (p *pointG1) Base() kyber.Point {
	//get a base
	panic("not implemented")
}

func (p *pointG1) Add(p1, p2 kyber.Point) kyber.Point {
	pg1 := p1.(*pointG1)
	pg2 := p2.(*pointG1)
	C.g1_add_w(&p.g, &pg1.g, &pg2.g)
	return p
}

func (p *pointG1) Sub(p1, p2 kyber.Point) kyber.Point {
	pg1 := p1.(*pointG1)
	pg2 := p2.(*pointG1)
	C.g1_sub_w(&p.g, &pg1.g, &pg2.g)
	return p
}

func (p *pointG1) Neg(p1 kyber.Point) kyber.Point {
	pg := p1.(*pointG1)
	C.g1_neg_w(&p.g, &pg.g)
	return p
}

func (p *pointG1) Mul(s kyber.Scalar, p1 kyber.Point) kyber.Point {
	//TODO: Waiting for a scalar implementation
	return p
}

func (p *pointG1) Clone() kyber.Point {
	p2 := new(pointG1)
	C.g1_new_w(&p2.g)
	C.g1_copy_w(&p2.g, &p.g)
	return p2
}

//******TODO******
func (p *pointG1) Data() ([]byte, error) {
	panic("not implemented")
}

func (p *pointG1) EmbedLen() int {
	panic("not implemented")
}

func (p *pointG1) Embed(data []byte, rand cipher.Stream) kyber.Point {
	panic("not implemented")
}

//***************

func (p *pointG1) Pick(rand cipher.Stream) kyber.Point {
	//TODO: rand is currently not being used because RELIC randomizes a point
	// by reference.
	C.g1_rand_w(&p.g)
	return p
}

func (p *pointG1) Set(p2 kyber.Point) kyber.Point {
	pg2 := p2.(*pointG1)
	C.g1_copy_w(&p.g, &pg2.g)
	return p
}

func (p *pointG1) SetVarTime(varTime bool) error {
	return errors.New("ErrVarTime")
}

func (p *pointG1) String() string {
	buff, err := p.MarshalBinary()
	if err != nil {
		panic("Error in marshalling")
	}
	return hex.EncodeToString(buff)
}

// -- -- -- -- Marshalling

func (p *pointG1) MarshalBinary() ([]byte, error) {
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
	C.g1_write_bin_w((*C.uint8_t)(b), clen, &p.g, pack)
	//We transform the unsafe.Pointer back to a []byte
	//GoBytes takes C data with explicity length and returns Go []byte
	copy(buff, C.GoBytes(b, clen))
	//TODO: Implement error-checking functionality
	return buff, nil
}

func (p *pointG1) MarshalTo(w io.Writer) (int, error) {
	return pointMarshalTo(p, w)
}

func (p *pointG1) UnmarshalBinary(buff []byte) error {
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
	C.g1_read_bin_w(&p.g, (*C.uint8_t)(b), len)
	//TODO: Implement error-checking functionality
	return nil
}

func (p *pointG1) UnmarshalFrom(r io.Reader) (int, error) {
	return pointUnmarshalFrom(p, r)
}

func (p *pointG1) MarshalSize() int {
	//Point compression flag
	pack := C.int(0)
	return int(C.g1_size_bin_w(&p.g, pack))
}

// -- -- -- -- Helper functions

func pointMarshalTo(p kyber.Point, w io.Writer) (int, error) {
	buf, err := p.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return w.Write(buf)
}

func pointUnmarshalFrom(p kyber.Point, r io.Reader) (int, error) {
	if strm, ok := r.(cipher.Stream); ok {
		p.Pick(strm)
		return -1, nil // no byte-count when picking randomly
	}
	buf := make([]byte, p.MarshalSize())
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return n, err
	}
	return n, p.UnmarshalBinary(buf)
}
