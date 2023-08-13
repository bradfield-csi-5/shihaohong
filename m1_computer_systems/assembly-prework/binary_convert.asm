section .text
global binary_convert
binary_convert: 			; rdi points to input string
	xor r8, r8 				; strlen

.getlen:
	movzx ecx, byte [rdi]	; each char is one byte
	cmp byte[rdi], 0
	je .start_convert

	inc r8
	add rdi, 1
	jmp .getlen

.start_convert:
sub rdi, r8 				; reset the addr of rdi
	xor rax, rax 			; init an accumulator (counter)

.get_char:
	movzx rcx, byte [rdi] 	; each char is one byte
	and rcx, 1 				; take only lowest order bit
	cmp rcx, 0
	je .check_char

	mov r9, 1 				; calc_register
	mov r10, 1 				; calc_iterator

.get_val:
	cmp r10, r8
	je .add
	imul r9, 2
	inc r10
	jmp .get_val

.add:
	add rax, r9 			; acc += calc_register

.check_char:
	dec r8
	add rdi, 1
	cmp r8, 0 				; got to end of the string
	jge .get_char

.end:
	ret
