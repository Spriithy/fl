#ifndef INTERNAL_COMMON_H
#define INTERNAL_COMMON_H

#include "memory.h"
#include "types.h"
#include <stdio.h>

/******************************************************************************/

#define fatalf(fmt, ...)                   \
    {                                      \
        fprintf(stderr, fmt, __VA_ARGS__); \
        exit(-1);                          \
    }

#define errorf(fmt, ...) fprintf(stderr, fmt, ...)

/******************************************************************************/

// silence some unexpected compiler warning
extern char* strdup(const char* s);

/******************************************************************************/

#endif /* common.h */