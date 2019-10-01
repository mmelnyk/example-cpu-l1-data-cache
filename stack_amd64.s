//+build amd64,!gccgo

#include "textflag.h"

TEXT ·asmStackIndex(SB),NOSPLIT,$-8-0
	MOVQ    SP, AX
    SHRQ    $10, AX
    MOVBLZX AL, AX
    MOVL    AX, stack+0(FP)
	RET
