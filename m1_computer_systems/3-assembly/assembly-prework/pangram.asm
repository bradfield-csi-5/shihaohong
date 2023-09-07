section .text
global pangram
pangram: ; rdi = string
	xor r8, r8 ; bit field for alphabet

check_char:
	movzx ecx, byte [rdi]
	cmp ecx, uppercase_low
	jl check_uppercase

check_lowercase: ; set that within range
	cmp ecx, uppercase_high
	jg check_uppercase ; not a lowercase character, jump to next char
	mov r9d, ecx
	sub r9d, uppercase_low
	bts r8d, r9d ; char is within range, set map
	jmp continue_to_next_char

check_uppercase:
	cmp ecx, lowercase_low
	jl continue_to_next_char
	cmp ecx, lowercase_high
	jg continue_to_next_char ; not a uppercase character, jump to next char
	mov r9d, ecx
	sub r9d, lowercase_low
	bts r8d, r9d ; char is within range, set map

continue_to_next_char:
	add rdi, 1
	cmp byte [rdi], 0
	jne check_char

check_map:
	xor rax, rax
	cmp r8, map_eq
	jne end
	mov rax, 1
end:
	ret

				section 	.rodata
map_eq			equ			67108863
lowercase_low	equ			65
lowercase_high	equ			90
uppercase_low	equ			97
uppercase_high	equ			122
