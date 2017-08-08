package main

// #include <relic.h>
/*
void g2_free_w(ep2_st* g2){
  g2_free(g2);
}
*/
import "C"

type ellPointG2 struct {
	g2 C.ep2_st
}

// ******* G2 METHODS *******

// // Initialization and Free

func freeG2(g2 *C.ep2_st) {
	C.g2_free_w(g2)
}
