// +build !race

// func XBegin(status int) (result int)
TEXT ·XBegin(SB),NOSPLIT,$8
	MOVQ	status+0(FP), AX
	BYTE $0xc7; BYTE $0xf8; BYTE $0x0; BYTE $0x0; BYTE $0x0; BYTE $0x0
	MOVQ	AX, result+8(FP)
	RET


