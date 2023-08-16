default rel

section .text
global volume
volume:
	mulss xmm0, xmm0 ; r * r
	mulss xmm0, xmm1 ; r^2 * h
	mulss xmm0, [pi] ; pi * r^2 * h
	cvtsi2ss xmm1, [const_3]
	divss xmm0, xmm1 ; pi * r^2 * h / 3
 	ret

			section 	.data
pi:			dd 			3.14159265359
const_3:	dd			3
