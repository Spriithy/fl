#ifndef INTERNAL_BYTECODE_H
#define INTERNAL_BYTECODE_H

/******************************************************************************/

typedef enum {
    Syscall,
    Save,
    Free,
    Load8,
    Load16,
    Load32,
    Load,
    Store8,
    Store16,
    Store32,
    Store,
    Const8,
    Const16,
    Const32,
    Const,
} v_bytecode;

/******************************************************************************/

#endif /* bytecode.h */