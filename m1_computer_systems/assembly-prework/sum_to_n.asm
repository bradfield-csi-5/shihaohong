section .text
global sum_to_n
sum_to_n:			; rdi is the input val
	xor ebx, ebx 	; reserve a register for total, set to zero
	xor r8d, r8d 		; set register for iterator

accumulate:
	add ebx, r8d
	inc r8d
	cmp r8d, edi
	jle accumulate

	mov eax, ebx 	; save total to rax
	ret
