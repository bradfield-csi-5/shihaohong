section .text
global fib
fib:
	cmp rdi, 1
	jg recurse
	mov rax, rdi
	ret

recurse:
	dec rdi
	push rdi
	call fib
	pop rdi
	push rax
	dec rdi
	call fib
	pop rdi
	add rax, rdi
	ret
