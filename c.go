package main

// #cgo LDFLAGS: -L/usr/local/lib/ -lrelic
// #cgo CFLAGS: -I/usr/local/include/relic
// #include <relic.h>
// #include <relic_test.h>
import "C"

/***
* relic-go is a Go wrapper for the RELIC library, with (currently) a focus solely on elliptic curves and its pairing crypto capabilities.
* The above declarations will need to be adjusted depending on your system's configuration (e.g. location of the RELIC library and headers).
* Also, several default configurations are set in RELIC's makefile. Make sure you take a close look at things such as security level and memory allocation.
* Thus far, this library has been written, configured, and tested with the factory defaults, on an OSX 64-bit machine.
* NOTE: It is very important to keep the following table as reference:

if FP_PRIME < 1536 (default is 256, from fp.cmake)
 G1_LOWER ep_
 G1_UPPER EP
 G2_LOWER ep2_
 G2_UPPER EP
 GT_LOWER fp12_
 PC_LOWER pp_
else
G1_LOWER ep_
G1_UPPER EP
G2_LOWER ep_
G2_UPPER EP
GT_LOWER fp2_
PC_LOWER pp_

* RELIC utilizes pre-processing directives to adapt its function calls accordingly. For example, g2_new() is a macro that will call either ep2_new() or ep_new() depending on FP_PRIME.
* The difference between them is that one implements a set of methods for elliptic curves over prime fields and the other over extensions of prime fields.
* Because Go doesn't have as much freedom to change a lot of things with pre-processing, it is VERY important to keep track of the type of functions that a group needs, according to the above configuration.
* Similarly, it is important to have in mind what MEMORY_ALLOC has been defined at compilation. The way ep_new/free behaves, for example, will vary accordingly and can be important to find bugs.
* Finally, the compilation architecture for RELIC matters a lot. Throughout the code, I assume the value of C int to be 32-bit and force the usage of int32. Casting types in such a way can become troublesome if this assumption is false.
***/
