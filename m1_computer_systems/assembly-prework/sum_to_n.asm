section .text
global sum_to_n
sum_to_n:			; rdi is the input val
	xor rbx, rbx 	; reserve a register for total, set to zero
	xor r8, r8 		; set register for iterator

accumulate:
	add rbx, r8
	inc r8
	cmp r8, rdi
	jle accumulate

	mov rax, rbx 	; save total to rax
	ret
