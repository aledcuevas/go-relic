package main

// #include <relic.h>
/*
void initialize_g1(ep_st* g1){
		TRY {
		g1_null(g1);
		g1_new(g1);
	} CATCH_ANY {
		//THROW(ERR_CAUGHT);
	} FINALLY {

	}
}
void g1_free_w(ep_st* g1){
  g1_free(g1);
}
*/
import "C"

type ellPointG1 struct {
	g1 C.ep_st
}

// ******* G1 METHODS *******

// // Initialization and Free
func freeG1(g1 *C.ep_st) {
	C.g1_free_w(g1)
}

func initG1(g1Strct *ellPointG1) {
	C.initialize_g1(&g1Strct.g1)
}
