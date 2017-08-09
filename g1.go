package main

// #include <relic.h>
/*
void g1_free_w(ep_st* g1){
  g1_free(g1);
}
void g1_null_w(ep_st* g1){
  g1_null(g1);
}
void g1_new_w(ep_st* g1){
  g1_new(g1);
}
*/
import "C"

// ******* G1 METHODS *******

type pointG1 struct {
	g1    C.ep_st
	lower int
	upper int
}

// // Initialization and Free

func freeG1(g1 *C.ep_st) {
	C.g1_free_w(g1)
}

func nullG1(g1 *C.ep_st) {
	C.g1_null_w(g1)
}

func randG1(g1 *C.ep_st) {
	C.ep_rand(g1)
}

func newG1(g1 *C.ep_st) {
	C.g1_new_w(g1)
}
