#include "stack.h"

/******************************************************************************/
v_stack* v_stack_create(size_t cap)
{
    v_stack* s = malloc(sizeof(*s));
    v_malloc_check(s, "v_stack.load()");

    s->cap = cap;
    s->size = 0;

    s->data = malloc(cap);
    v_malloc_check(s->data, "v_stack.load()");
    s->t8 = s->data;

    return s;
}

/******************************************************************************/
void v_stack_destroy(v_stack* s)
{
    if (!s)
        return;

    v_free(s->data);
    v_free(s);
}

/******************************************************************************/
void push_u64(v_stack* s, u64 v)
{
    if (s->size + 8 < s->cap) {
        *s->t64++ = v;
        s->size += 8;
        return;
    }

    fatalf("v_stack.push(%p, u64): stack overflow\n", s);
}

/******************************************************************************/
void push_u32(v_stack* s, u32 v)
{
    if (s->size + 4 < s->cap) {
        *s->t32++ = v;
        s->size += 4;
        return;
    }

    fatalf("v_stack.push(%p, u32): stack overflow\n", s);
}

/******************************************************************************/
void push_u16(v_stack* s, u16 v)
{
    if (s->size + 2 < s->cap) {
        *s->t16++ = v;
        s->size += 2;
        return;
    }

    fatalf("v_stack.push(%p, u16): stack overflow\n", s);
}

/******************************************************************************/
void push_u8(v_stack* s, u8 v)
{
    if (s->size + 1 < s->cap) {
        *s->t8++ = v;
        s->size++;
        return;
    }

    fatalf("v_stack.push(%p, u8): stack overflow\n", s);
}

/******************************************************************************/
u64 pop_u64(v_stack* s)
{
    if (s->size < 8)
        fatalf("v_stack.pop(%p) -> u64: stack underflow\n", s);

    s->size -= 8;
    return *--s->t64;
}

/******************************************************************************/
u32 pop_u32(v_stack* s)
{
    if (s->size < 4)
        fatalf("v_stack.pop(%p) -> u32: stack underflow\n", s);

    s->size -= 4;
    return *--s->t32;
}

/******************************************************************************/
u16 pop_u16(v_stack* s)
{
    if (s->size < 2)
        fatalf("v_stack.pop(%p) -> u16: stack underflow\n", s);

    s->size -= 2;
    return *--s->t16;
}

/******************************************************************************/
u8 pop_u8(v_stack* s)
{
    if (s->size < 1)
        fatalf("v_stack.pop(%p) -> u8: stack underflow\n", s);

    s->size--;
    return *--s->t8;
}

/******************************************************************************/
void v_stack_dump(v_stack* s)
{
    if (!s)
        fatalf("v_stack.dump(%p): null v_stack reference\n", s);

    puts("~ v_stack dump ~");
    printf("~ cap : %zu\n", s->cap);
    printf("~ len : %zu\n", s->size);
    printf("~  sp : %p (0x%08x)\n", s->t8, *(s->t32 - 1));
}