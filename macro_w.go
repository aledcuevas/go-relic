package main

// #include <relic.h>
/*
void pc_param_print_w(){
	pc_param_print();
}
*/
import "C"

/***
 * RELIC macro wrappers.
 ***/

func paramPrint() {
	C.pc_param_print_w()
}
