package gotsx

// Package gotsx provides low-level access to TSX primitives
// on Intel machines that have TSX-NI extensions.
//
// These may be useful for implementing low-level concurrent algorithms.
//
// These functions require great care to be used correctly.

// Helper constants.
const (
	XA_EXPLICIT = 0
	XA_RETRY    = 1
	XA_CONFLICT = 2
	XA_CAPACITY = 3

	XBEGIN_STARTED  = ^0
	XABORT_EXPLICIT = (1 << XA_EXPLICIT)
	XABORT_RETRY    = (1 << XA_RETRY)
	XABORT_CONFLICT = (1 << XA_CONFLICT)
	XABORT_CAPACITY = (1 << XA_CAPACITY)
	XABORT_DEBUG    = (1 << 4)
	XABORT_NESTED   = (1 << 5)
)

// XBegin starts transactional execution. It is expected that XBEGIN_STARTED is passed as status. If 
// the transaction starts succesfully the user can compare result against XBEGIN_STARTED, they should
// be equal, if not, result will contain the abort status of the transaction.
func XBegin(status int) (result int)

// XEnd ends transactional execution, it attempts to commit the transactions speculative working set.
func XEnd()

// XTest can be used to check if the code being executed is in a transaction. result will be zero if
// the code is not being executed transactionally, and non-zero if it is.
func XTest() (result int)

// Need to work out how to pass an arbitrary imm8 to an assembly function.
// func XAbort(status uint8) (result int)

// XAbort2 explicitly aborts a transaction. It writes the abort status 0x2 into the transaction's 
// status register, which set the upper bits of result.
func XAbort2() (result int)




