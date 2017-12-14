#ifndef INTERNAL_TYPES_H
#define INTERNAL_TYPES_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

typedef uint8_t u8;
typedef uint16_t u16;
typedef uint32_t u32;
typedef uint64_t u64;

typedef int8_t i8;
typedef int16_t i16;
typedef int32_t i32;
typedef int64_t i64;

typedef float f32;
typedef double f64;

/******************************************************************************/

u64 f64_bits(f64 v);
f64 f64_from_bits(u64 v);

u32 f32_bits(f32 v);
f32 f32_from_bits(u32 v);

#endif /* types.h */