package main

// #cgo LDFLAGS: -L/usr/local/lib/ -lrelic
// #cgo CFLAGS: -I/usr/local/include/relic
// #include <relic.h>
// #include <relic_test.h>
import "C"
