section .text
global leap_year

leap_year:
    mov ebx, edi
    and ebx, 3 ; check if mod 4
    cmp ebx, 0
    jnz not_leap

    xor edx, edx
    mov eax, edi
    mov ecx, 100
    div ecx
    cmp rdx, 0 ; check if mod 100
    jnz leap

    xor edx, edx
    mov eax, edi
    mov ecx, 400
    div ecx ;
    cmp edx, 0 ; check if mod 400
    jz leap

not_leap:
    xor rax, rax
    jmp done
leap:
    mov rax, 1
done:
    ret
