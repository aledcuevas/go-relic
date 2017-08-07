package main

// #include<relic.h>
import "C"

/***
 * List of useful constants.
 * Msgs found in relic_err.h
 * RELIC constants in relic_core.h
 ***/

const (
	//TRUE const to check against C functions that return int
	TRUE = 0
	//FALSE const to check against C functions that return int
	FALSE = 1

	// MsgNoMemory Error message respective to ERR_NO_MEMORY.
	MsgNoMemory = "not enough memory"
	// MsgNoPreci Error message respective to ERR_PRECISION.
	MsgNoPreci = "insufficient precision"
	// MsgNoFile Error message respective to ERR_NO FILE.
	MsgNoFile = "file not found"
	// MsgNoRead Error message respective to ERR_NO_READ.
	MsgNoRead = "can't read bytes from file"
	// MsgNoValid Error message respective to ERR_NO_VALID.
	MsgNoValid = "invalid value passed as input"
	// MsgNoBuffer Error message respective to ERR_NO_BUFFER.
	MsgNoBuffer = "insufficient buffer capacity"
	// MsgNoField Error message respective to ERR_NO_FIELD.
	MsgNoField = "no finite field supported at this security level"
	// MsgNoCurve Error message respective to ERR_NO_CURVE.
	MsgNoCurve = "no curve supported at this security level"
	// MsgNoConfig Error message respective to ERR_NO_CONFIG.
	MsgNoConfig = "invalid library configuration"

	// StsOk Indicates that the function executed correctly.
	StsOk = 0
	// StsErr Indicates that an error occurred during the function execution.
	StsErr = 1
	// CmpLt Indicates that a comparison returned that the first argument was lesser than the second argument.
	CmpLt = -1
	// CmpEq Indicates that a comparison returned that the first argument was equal to the second argument.
	CmpEq = 0
	// CmpGt Indicates that a comparison returned that the first argument was greater than the second argument.
	CmpGt = 1
	// CmpNe Indicates that two incomparable elements are not equal.
	CmpNe = 2
	// OptZero Optimization identifer for the case where a coefficient is 0.
	OptZero = 0
	// OptOne Optimization identifer for the case where a coefficient is 1.
	OptOne = 1
	// OptTwo Optimization identifer for the case where a coefficient is 1.
	OptTwo = 2
	// OptDigit Optimization identifer for the case where a coefficient is small.
	OptDigit = 3
	// OptMinus3 Optimization identifier for the case where a coefficient is -3.
	OptMinus3 = 4
	// RelicOptNone Optimization identifier for the case where the coefficient is random
	RelicOptNone = 5
	// MaxTerms Maximum number of terms to describe a sparse object.
	MaxTerms = 16

	// PcBytes (TEMPORARY) represents the size of bytes of a block sufficient to store a binary field element. Will later be migrated to goCore
	PcBytes = int(C.FP_BYTES)
)
