package main

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

func XBegin(status int) (result int)
func XEnd()
func XTest() (result int)

// Need to work out how to pass an arbitrary imm8 to an assembly function.
// func XAbort(status uint8) (result int)

// For now, I've just implemented an xabort with 2 as the immediate.
func XAbort2() (result int)




