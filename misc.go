package main

// #include<relic.h>
import "C"

const (
	//TRUE const to check against C functions that return int
	TRUE = 0
	//FALSE const to check against C functions that return int
	FALSE = 1

	// MsgNoMemory Error message respective to ERR_NO_MEMORY. */
	MsgNoMemory = "not enough memory"
	// MsgNoPreci Error message respective to ERR_PRECISION. */
	MsgNoPreci = "insufficient precision"
	// MsgNoFile Error message respective to ERR_NO FILE. */
	MsgNoFile = "file not found"
	// MsgNoRead Error message respective to ERR_NO_READ. */
	MsgNoRead = "can't read bytes from file"
	// MsgNoValid Error message respective to ERR_NO_VALID. */
	MsgNoValid = "invalid value passed as input"
	// MsgNoBuffer Error message respective to ERR_NO_BUFFER. */
	MsgNoBuffer = "insufficient buffer capacity"
	// MsgNoField Error message respective to ERR_NO_FIELD. */
	MsgNoField = "no finite field supported at this security level"
	// MsgNoCurve Error message respective to ERR_NO_CURVE. */
	MsgNoCurve = "no curve supported at this security level"
	// MsgNoConfig Error message respective to ERR_NO_CONFIG. */
	MsgNoConfig = "invalid library configuration"
)
