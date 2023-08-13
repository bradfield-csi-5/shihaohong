section .text
global pangram
pangram: ; rdi = string
	xor r8, r8 ; bit field for alphabet

check_char:
	movzx ecx, byte [rdi]
	cmp ecx, 97
	jge is_lowercase_lower

continue_to_next_char:
	add rdi, 1
	cmp byte [rdi], 0
	je check_map
	jmp check_char

is_lowercase_lower: ; set that within range
	cmp ecx, 122
	jg continue_to_next_char ; not a lowercase character, jump to next char

is_lowercase_upper: ; char is within range, set map
	mov r9d, ecx
	sub r9d, 97
	bts r8d, r9d

	jmp continue_to_next_char

; ignore upper case for now

check_map:
	xor rax, rax
	cmp r8, mapEq
	jne end
	mov rax, 1
end:
	ret


		section 	.bss
mapEq	equ			67108863