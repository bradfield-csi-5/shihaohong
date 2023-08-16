section .text
global binary_convert
binary_convert:
	xor r8, r8 ; strlen

.getstrlen:
	movzx ecx, byte [rdi]
	cmp byte[rdi], 0 ; logic to cover strlen = 0
	je .start_convert

	inc r8
	add rdi, 1
	jmp .getstrlen

.start_convert:
	sub rdi, r8 ; reset rdi to initial position
	xor rax, rax ; init ret value to 0

.get_char:
	movzx rcx, byte [rdi]
	and rcx, 1
	cmp rcx, 0
	je .check_next_char ; if bit is 0, skip calculations

	mov r9, 1 ; init mult init value
	mov r10, 1 ; init mult iterator

.exp_2:
	cmp r10, r8
	je .accumulate
	imul r9, 2
	inc r10
	jmp .exp_2

.accumulate:
	add rax, r9

.check_next_char:
	dec r8
	add rdi, 1
	cmp r8, 0
	jge .get_char

.end:
	ret
