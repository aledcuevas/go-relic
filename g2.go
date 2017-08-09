package main

// #include <relic.h>
/*
void g2_free_w(ep2_st* g2){
  g2_free(g2);
}
void g2_null_w(ep2_st* g2){
  g2_null(g2);
}
void g2_new_w(ep2_st* g2){
  g2_new(g2);
}
*/
import "C"

// ******* G2 METHODS *******

type pointG2 struct {
	g2    C.ep2_st
	lower int
	upper int
}

// // Initialization and Free

func freeG2(g2 *C.ep2_st) {
	C.g2_free_w(g2)
}

func nullG2(g2 *C.ep2_st) {
	C.g2_null_w(g2)
}

func randG2(g2 *C.ep2_st) {
	C.ep2_rand(g2)
}

func newG2(g2 *C.ep2_st) {
	C.g2_new_w(g2)
}
