package main

// #include <relic.h>
/*
void initialize_gt(fp12_t* gt){
		TRY {
		gt_null(gt);
		gt_new(gt);
	} CATCH_ANY {
		//THROW(ERR_CAUGHT);
	} FINALLY {

	}
}
void gt_free_w(fp12_t* gt){
  gt_free(gt);
}
*/
import "C"

type ellPointGT struct {
	gt C.fp12_t
}

// ******* GT METHODS *******

// // Initialization and Free

func freeGT(gt *C.fp12_t) {
	C.gt_free_w(gt)
}

func initGT(gtStrct *ellPointGT) {
	C.initialize_gt(&gtStrct.gt)
}
