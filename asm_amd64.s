// +build !race

// func XBegin(status int) (result int)
TEXT ·XBegin(SB),$0-16
	MOVQ	status+0(FP), AX
	BYTE $0xc7; BYTE $0xf8; BYTE $0x0; BYTE $0x0; BYTE $0x0; BYTE $0x0
	MOVQ	AX, result+8(FP)
	RET

TEXT ·XEnd(SB),$0-0
	BYTE $0x0f; BYTE $0x01; BYTE $0xd5
	RET

TEXT ·XTest(SB),$0-8
	BYTE $0x0f; BYTE $0x01; BYTE $0xd6;
	SETNE   AX
	MOVQ	AX, result+0(FP)
	RET

TEXT ·XAbort2(SB),$0-8
	BYTE $0xc6; BYTE $0xf8; BYTE $0x2;
	MOVQ	AX, result+0(FP)
	RET


