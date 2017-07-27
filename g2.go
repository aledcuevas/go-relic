package main

// #include <relic.h>
/*
void initialize_g2(ep2_st* g2){
		TRY {
		g2_null(g2);
		g2_new(g2);
	} CATCH_ANY {
		//THROW(ERR_CAUGHT);
	} FINALLY {

	}
}
void g2_free_w(ep2_st* g2){
  g2_free(g2);
}
*/
import "C"

type ellPointG2 struct {
	g2 C.ep2_st
}

// ******* G1 METHODS *******

// // Initialization and Free

func freeG2(g2 *C.ep2_st) {
	C.g2_free_w(g2)
}

func initG2(g2Strct *ellPointG2) {
	C.initialize_g2(&g2Strct.g2)
}
