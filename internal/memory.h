#ifndef INTERNAL_MEMORY_H
#define INTERNAL_MEMORY_H

#include "types.h"
#include <endian.h>
#include <stdio.h>
#include <stdlib.h>

/******************************************************************************/

#define v_malloc_check(ptr, msg) \
    if (!(ptr)) {                \
        perror(msg);             \
        exit(-1);                \
    }

#define v_free(ptr)   \
    if (ptr) {        \
        free(ptr);    \
        (ptr) = NULL; \
    }

/******************************************************************************/

#define __gen_data_u8__(v) v & 0xff

#if __BYTE_ORDER == __LITTLE_ENDIAN
#define __gen_data_u16__(v) __gen_data_u8__(v), __gen_data_u8__(v >> 8)
#define __gen_data_u32__(v) __gen_data_u16__(v), __gen_data_u16__(v >> 16)
#define __gen_data_u64__(v) __gen_data_u32__(v), __gen_data_u32__(v >> 32)
#elif __BYTE_ORDER == __BIG_ENDIAN
#define __gen_data_u16__(v) __gen_data_u8__(v >> 8), __gen_data_u8__(v)
#define __gen_data_u32__(v) __gen_data_u16__(v >> 16), __gen_data_u16__(v)
#define __gen_data_u64__(v) __gen_data_u32__(v >> 32), __gen_data_u32__(v)
#else
#error "unsupported endianness"
#endif

#define __gen_data_f64__(v) __gen_data_u64__(f64_bits(v))
#define __gen_data_f32__(v) __gen_data_u32__(f32_bits(v))

#define __gen_data_i64__(v) __gen_data_u64__((u64)(v))
#define __gen_data_i32__(v) __gen_data_u32__((u32)(v))
#define __gen_data_i16__(v) __gen_data_u16__((u16)(v))
#define __gen_data_i8__(v) __gen_data_u8__((u8)(v))

/******************************************************************************/

#endif /* memory.h */