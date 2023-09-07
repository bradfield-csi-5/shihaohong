section .text
global index
index:
	; rdi: matrix
	; rsi: rows
	; rdx: cols
	; rcx: rindex
	; r8: cindex

	; calculate the number of rows to jump
	; need to determine size of row
	mov r9, 1
	imul r9, rcx
	imul r9, rdx
	lea rdi, [rdi + 4 * r9] ; handle row jump
	mov rax, [rdi + 4 * r8] ; handle column jump

	ret
