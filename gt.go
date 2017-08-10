package main

// #include <relic.h>
/*
void gt_free_w(gt_t* p){
  gt_free(p);
}
void gt_null_w(gt_t* p){
  gt_null(p);
}
void gt_new_w(gt_t* p){
  gt_new(p);
}
void gt_rand_w(gt_t p){
	//BUG: For some reason the signature is not being checked properly here
	// Temporarily removed pointers to continue testing
  gt_rand(p);
}
int gt_cmp_w(gt_t p, gt_t q){
	//BUG: For some reason the signature is not being checked properly here
	// Temporarily removed pointers to continue testing
  return gt_cmp(p,q);
}
*/
import "C"
import kyber "gopkg.in/dedis/kyber.v1"

type pointGT struct {
	g C.gt_t
	p *Pairing
}

func newPointGT(p *Pairing) *pointGT {
	pg := new(pointGT)
	C.gt_new_w(&pg.g)
	//runtime.SetFinalizer(&pg.g, clear)
	return pg
}

func (p *pointGT) Pairing(p1, p2 kyber.Point) kyber.Point {
	panic("missing implementation")
}

func (p *pointGT) Equal(p2 kyber.Point) bool {
	pg := p2.(*pointGT)
	i := C.gt_cmp_w(&p.g, &pg.g)
	switch i {
	case C.CMP_EQ:
		return true
	case C.CMP_NE:
		return false
	default:
		panic("Error in C casting")
	}
}

func (p *pointGT) Null() kyber.Point {
	C.gt_null_w(&p.g)
	return p
}
