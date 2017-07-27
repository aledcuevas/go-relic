package main

// #include <relic.h>
/*
void pc_param_print_w(){
	pc_param_print();
}
*/
import "C"

func paramPrint() {
	C.pc_param_print_w()
}
