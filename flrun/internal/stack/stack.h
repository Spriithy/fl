#ifndef INTERNAL_STACK_H
#define INTERNAL_STACK_H

#include "../common.h"

/******************************************************************************/

typedef struct {
    size_t cap;
    size_t size;
    u8*    data;
    union {
        u64* t64;
        u32* t32;
        u16* t16;
        u8*  t8;
    };
} fl_stack;

fl_stack* fl_stack_create(size_t cap);
void      fl_stack_dump(fl_stack* s);
void      fl_stack_destroy(fl_stack* s);

void push_u64(fl_stack* s, u64 v);
void push_u32(fl_stack* s, u32 v);
void push_u16(fl_stack* s, u16 v);
void push_u8(fl_stack* s, u8 v);

u64 pop_u64(fl_stack* s);
u32 pop_u32(fl_stack* s);
u16 pop_u16(fl_stack* s);
u8  pop_u8(fl_stack* s);

/******************************************************************************/

#define __float_to_unsigned__(v, sz) *(u##sz*)(&v)
#define __unsigned_to_float__(v, sz) *(f##sz*)(&v)

#define push_f64(s, v) push_u32(s, __float_to_unsigned__(v, 64))
#define push_f32(s, v) push_u32(s, __float_to_unsigned__(v, 32))

#define pop_f64(s) __unsigned_to_float__(pop_u64(s), 64)
#define pop_f32(s) __unsigned_to_float__(pop_u32(s), 32)

#define push(s, v) _Generic((v), u64        \
                            : push_u64, u32 \
                            : push_u32, u16 \
                            : push_u16, u8  \
                            : push_u8, f64  \
                            : push_f64, f32 \
                            : push_f32)(s, v)

#endif /* fl_stack.h */