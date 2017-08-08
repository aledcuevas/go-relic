package main

// #include <relic.h>
/*
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
