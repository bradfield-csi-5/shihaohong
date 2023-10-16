#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct DoublyLinkedList DoublyLinkedList;
struct DoublyLinkedList {
    DoublyLinkedList* prev;
    DoublyLinkedList* next;
    char* val;
};

DoublyLinkedList* newDLLNode(char* val) {
    DoublyLinkedList* dll = malloc(sizeof(DoublyLinkedList));
    dll->next = NULL;
    dll->prev = NULL;
    dll->val = val;
    return dll;
}

DoublyLinkedList* insert(DoublyLinkedList* dll, char* val) {
    while (dll->next != NULL) {
        dll = dll->next;
    }

    DoublyLinkedList* newNode = newDLLNode(val);
    dll->next = newNode;
    newNode->prev = dll;
    return newNode;
}

DoublyLinkedList* find(DoublyLinkedList* dll, char* val) {
    while ((dll->next != NULL)) {
        if (strcmp(dll->val, val) == 0) {
            return dll;
        }

        dll = dll->next;
    }

    if (strcmp(dll->val, val) == 0) {
        return dll;
    }

    return NULL;
}

// -1 not found, 0 found and removed
// cannot delete last node
int delete(DoublyLinkedList* dll, char* val) {
    while ((dll->next != NULL)) {
        if (strcmp(dll->val, val) == 0) {
            if (dll->next != NULL) {
                dll->prev->next = dll->next;
                dll->next->prev = dll->prev;
            } else {
                dll->prev->next = NULL;
            }
            free(dll);
            return 0;
        }

        dll = dll->next;
    }

    if (strcmp(dll->val, val) == 0) {
            if (dll->next != NULL) {
                dll->prev->next = dll->next;
                dll->next->prev = dll->prev;
            } else {
                dll->prev->next = NULL;
            }
            free(dll);
            return 0;
    }

    return -1;
}

int main() {
    // test find
    // DoublyLinkedList* first = newDLLNode("first node");
    // DoublyLinkedList* result = find(first, "second node");

    // printf("first node prev: %p\n", first->prev);
    // printf("first node next: %p\n", first->next);
    // printf("first node s: %s\n", first->val);
    // printf("RESULT: %p\n", result);

    // DoublyLinkedList* second = insert(first, "second node");
    // result = find(first, "second node");

    // printf("first node prev: %p\n", first->prev);
    // printf("first node next: %p\n", first->next);
    // printf("first node s: %s\n", first->val);

    // printf("second prev: %p\n", second->prev);
    // printf("second next: %p\n", second->next);
    // printf("second s: %s\n", second->val);

    // printf("RESULT: %p\n", result);
    // printf("RESULT val: %s\n", result->val);

    // test create and delete
    DoublyLinkedList* first = newDLLNode("first node");
    DoublyLinkedList* second = insert(first, "second node");
    DoublyLinkedList* third = insert(first, "third node");
    DoublyLinkedList* fourth = insert(first, "fourth node");
    DoublyLinkedList* fifth = insert(first, "fifth node");

    // third
    printf("BEFORE DELETE\n");
    printf("second node prev: %p\n", second->prev);
    printf("second node next: %p\n", second->next);
    printf("third node val from second: %s\n", second->next->val);
    printf("third node val from fourth: %s\n", second->next->val);

    delete(first, "third node");
    printf("AFTER DELETE\n");
    printf("third node val from second: %s\n", second->next->val);
    printf("third node val from fourth: %s\n", fourth->prev->val);

    // last
    printf("BEFORE DELETE\n");
    printf("fourth node next: %p\n", fourth->next);
    printf("fifth val from fourth: %s\n", fourth->next->val);

    delete(first, "fifth node");
    printf("AFTER DELETE\n");
    printf("fourth node next: %p\n", fourth->next);
    printf("fifth: %p\n", fifth);
}
