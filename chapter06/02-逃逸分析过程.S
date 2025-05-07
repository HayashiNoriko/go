main.foo STEXT size=261 args=0x8 locals=0x48 funcid=0x0 align=0x0
	0x0000 00000 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	TEXT	main.foo(SB), ABIInternal, $72-8
	0x0000 00000 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	CMPQ	SP, 16(R14)
	0x0004 00004 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	PCDATA	$0, $-2
	0x0004 00004 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	JLS	240
	0x000a 00010 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	PCDATA	$0, $-1
	0x000a 00010 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	PUSHQ	BP
	0x000b 00011 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	MOVQ	SP, BP
	0x000e 00014 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	SUBQ	$64, SP
	0x0012 00018 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	FUNCDATA	$0, gclocals·ISb46fRPFoZ9pIfykFK/kQ==(SB)
	0x0012 00018 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	FUNCDATA	$1, gclocals·jCgrU8XAg0ifiSJZPFgpKw==(SB)
	0x0012 00018 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	FUNCDATA	$5, main.foo.arginfo1(SB)
	0x0012 00018 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	MOVQ	AX, main.arg_val+80(SP)
	0x0017 00023 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:9)	MOVQ	$11, main.foo_val1+48(SP)
	0x0020 00032 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:10)	MOVQ	$22, main.foo_val2+40(SP)
	0x0029 00041 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:11)	LEAQ	type:int(SB), AX
	0x0030 00048 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:11)	PCDATA	$1, $0
	0x0030 00048 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:11)	CALL	runtime.newobject(SB)
	0x0035 00053 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:11)	MOVQ	AX, main.&foo_val3+56(SP)
	0x003a 00058 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:11)	MOVQ	$33, (AX)
	0x0041 00065 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:12)	MOVQ	$44, main.foo_val4+32(SP)
	0x004a 00074 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:13)	MOVQ	$55, main.foo_val5+24(SP)
	0x0053 00083 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:13)	XORL	CX, CX
	0x0055 00085 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:17)	JMP	224
	0x005a 00090 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:17)	MOVQ	CX, main.i+16(SP)
	0x005f 00095 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	PCDATA	$1, $1
	0x005f 00095 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	NOP
	0x0060 00096 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printlock(SB)
	0x0065 00101 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	LEAQ	main.arg_val+80(SP), AX
	0x006a 00106 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printpointer(SB)
	0x006f 00111 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printsp(SB)
	0x0074 00116 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	LEAQ	main.foo_val1+48(SP), AX
	0x0079 00121 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printpointer(SB)
	0x007e 00126 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	NOP
	0x0080 00128 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printsp(SB)
	0x0085 00133 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	LEAQ	main.foo_val2+40(SP), AX
	0x008a 00138 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printpointer(SB)
	0x008f 00143 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printsp(SB)
	0x0094 00148 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	MOVQ	main.&foo_val3+56(SP), AX
	0x0099 00153 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printpointer(SB)
	0x009e 00158 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	NOP
	0x00a0 00160 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printsp(SB)
	0x00a5 00165 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	LEAQ	main.foo_val4+32(SP), AX
	0x00aa 00170 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printpointer(SB)
	0x00af 00175 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printsp(SB)
	0x00b4 00180 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	LEAQ	main.foo_val5+24(SP), AX
	0x00b9 00185 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printpointer(SB)
	0x00be 00190 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	NOP
	0x00c0 00192 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printnl(SB)
	0x00c5 00197 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:18)	CALL	runtime.printunlock(SB)
	0x00ca 00202 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:17)	MOVQ	main.i+16(SP), CX
	0x00cf 00207 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:17)	INCQ	CX
	0x00d2 00210 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:22)	MOVQ	main.&foo_val3+56(SP), AX
	0x00d7 00215 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:22)	NOP
	0x00e0 00224 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:17)	CMPQ	CX, $5
	0x00e4 00228 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:17)	JLT	90
	0x00ea 00234 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:22)	ADDQ	$64, SP
	0x00ee 00238 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:22)	POPQ	BP
	0x00ef 00239 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:22)	RET
	0x00f0 00240 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:22)	NOP
	0x00f0 00240 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	PCDATA	$1, $-1
	0x00f0 00240 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	PCDATA	$0, $-2
	0x00f0 00240 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	MOVQ	AX, 8(SP)
	0x00f5 00245 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x00fa 00250 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	PCDATA	$0, $-1
	0x00fa 00250 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	MOVQ	8(SP), AX
	0x00ff 00255 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	NOP
	0x0100 00256 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:8)	JMP	0
	0x0000 49 3b 66 10 0f 86 e6 00 00 00 55 48 89 e5 48 83  I;f.......UH..H.
	0x0010 ec 40 48 89 44 24 50 48 c7 44 24 30 0b 00 00 00  .@H.D$PH.D$0....
	0x0020 48 c7 44 24 28 16 00 00 00 48 8d 05 00 00 00 00  H.D$(....H......
	0x0030 e8 00 00 00 00 48 89 44 24 38 48 c7 00 21 00 00  .....H.D$8H..!..
	0x0040 00 48 c7 44 24 20 2c 00 00 00 48 c7 44 24 18 37  .H.D$ ,...H.D$.7
	0x0050 00 00 00 31 c9 e9 86 00 00 00 48 89 4c 24 10 90  ...1......H.L$..
	0x0060 e8 00 00 00 00 48 8d 44 24 50 e8 00 00 00 00 e8  .....H.D$P......
	0x0070 00 00 00 00 48 8d 44 24 30 e8 00 00 00 00 66 90  ....H.D$0.....f.
	0x0080 e8 00 00 00 00 48 8d 44 24 28 e8 00 00 00 00 e8  .....H.D$(......
	0x0090 00 00 00 00 48 8b 44 24 38 e8 00 00 00 00 66 90  ....H.D$8.....f.
	0x00a0 e8 00 00 00 00 48 8d 44 24 20 e8 00 00 00 00 e8  .....H.D$ ......
	0x00b0 00 00 00 00 48 8d 44 24 18 e8 00 00 00 00 66 90  ....H.D$......f.
	0x00c0 e8 00 00 00 00 e8 00 00 00 00 48 8b 4c 24 10 48  ..........H.L$.H
	0x00d0 ff c1 48 8b 44 24 38 66 0f 1f 84 00 00 00 00 00  ..H.D$8f........
	0x00e0 48 83 f9 05 0f 8c 70 ff ff ff 48 83 c4 40 5d c3  H.....p...H..@].
	0x00f0 48 89 44 24 08 e8 00 00 00 00 48 8b 44 24 08 90  H.D$......H.D$..
	0x0100 e9 fb fe ff ff                                   .....
	rel 44+4 t=R_PCREL type:int+0
	rel 49+4 t=R_CALL runtime.newobject+0
	rel 97+4 t=R_CALL runtime.printlock+0
	rel 107+4 t=R_CALL runtime.printpointer+0
	rel 112+4 t=R_CALL runtime.printsp+0
	rel 122+4 t=R_CALL runtime.printpointer+0
	rel 129+4 t=R_CALL runtime.printsp+0
	rel 139+4 t=R_CALL runtime.printpointer+0
	rel 144+4 t=R_CALL runtime.printsp+0
	rel 154+4 t=R_CALL runtime.printpointer+0
	rel 161+4 t=R_CALL runtime.printsp+0
	rel 171+4 t=R_CALL runtime.printpointer+0
	rel 176+4 t=R_CALL runtime.printsp+0
	rel 186+4 t=R_CALL runtime.printpointer+0
	rel 193+4 t=R_CALL runtime.printnl+0
	rel 198+4 t=R_CALL runtime.printunlock+0
	rel 246+4 t=R_CALL runtime.morestack_noctxt+0
main.main STEXT size=92 args=0x0 locals=0x20 funcid=0x0 align=0x0
	0x0000 00000 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	TEXT	main.main(SB), ABIInternal, $32-0
	0x0000 00000 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	CMPQ	SP, 16(R14)
	0x0004 00004 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	PCDATA	$0, $-2
	0x0004 00004 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	JLS	85
	0x0006 00006 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	PCDATA	$0, $-1
	0x0006 00006 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	PUSHQ	BP
	0x0007 00007 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	MOVQ	SP, BP
	0x000a 00010 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	SUBQ	$24, SP
	0x000e 00014 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	FUNCDATA	$0, gclocals·ISb46fRPFoZ9pIfykFK/kQ==(SB)
	0x000e 00014 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	FUNCDATA	$1, gclocals·jCgrU8XAg0ifiSJZPFgpKw==(SB)
	0x000e 00014 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:26)	MOVL	$666, AX
	0x0013 00019 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:26)	PCDATA	$1, $0
	0x0013 00019 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:26)	CALL	main.foo(SB)
	0x0018 00024 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:26)	MOVQ	AX, main.main_val+16(SP)
	0x001d 00029 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	MOVQ	(AX), CX
	0x0020 00032 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	MOVQ	CX, main..autotmp_2+8(SP)
	0x0025 00037 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	PCDATA	$1, $1
	0x0025 00037 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	CALL	runtime.printlock(SB)
	0x002a 00042 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	MOVQ	main..autotmp_2+8(SP), AX
	0x002f 00047 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	CALL	runtime.printint(SB)
	0x0034 00052 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	CALL	runtime.printsp(SB)
	0x0039 00057 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	MOVQ	main.main_val+16(SP), AX
	0x003e 00062 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	PCDATA	$1, $0
	0x003e 00062 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	NOP
	0x0040 00064 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	CALL	runtime.printpointer(SB)
	0x0045 00069 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	CALL	runtime.printnl(SB)
	0x004a 00074 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:27)	CALL	runtime.printunlock(SB)
	0x004f 00079 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:28)	ADDQ	$24, SP
	0x0053 00083 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:28)	POPQ	BP
	0x0054 00084 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:28)	RET
	0x0055 00085 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:28)	NOP
	0x0055 00085 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	PCDATA	$1, $-1
	0x0055 00085 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	PCDATA	$0, $-2
	0x0055 00085 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	CALL	runtime.morestack_noctxt(SB)
	0x005a 00090 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	PCDATA	$0, $-1
	0x005a 00090 (/Users/mengmeng/go/src/hello/chapter06/02-逃逸分析过程.go:25)	JMP	0
	0x0000 49 3b 66 10 76 4f 55 48 89 e5 48 83 ec 18 b8 9a  I;f.vOUH..H.....
	0x0010 02 00 00 e8 00 00 00 00 48 89 44 24 10 48 8b 08  ........H.D$.H..
	0x0020 48 89 4c 24 08 e8 00 00 00 00 48 8b 44 24 08 e8  H.L$......H.D$..
	0x0030 00 00 00 00 e8 00 00 00 00 48 8b 44 24 10 66 90  .........H.D$.f.
	0x0040 e8 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48  ...............H
	0x0050 83 c4 18 5d c3 e8 00 00 00 00 eb a4              ...]........
	rel 20+4 t=R_CALL main.foo+0
	rel 38+4 t=R_CALL runtime.printlock+0
	rel 48+4 t=R_CALL runtime.printint+0
	rel 53+4 t=R_CALL runtime.printsp+0
	rel 65+4 t=R_CALL runtime.printpointer+0
	rel 70+4 t=R_CALL runtime.printnl+0
	rel 75+4 t=R_CALL runtime.printunlock+0
	rel 86+4 t=R_CALL runtime.morestack_noctxt+0
go:cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
main..inittask SNOPTRDATA size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
gclocals·ISb46fRPFoZ9pIfykFK/kQ== SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·jCgrU8XAg0ifiSJZPFgpKw== SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
main.foo.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
