#ifndef EXEC_VM_H
#define EXEC_VM_H

#include "../internal/bytecode.h"
#include "../internal/common.h"
#include "../internal/stack/stack.h"

/******************************************************************************/

typedef struct {
    fl_stack*    stack;
    size_t       pc;
    fl_bytecode* code;
    size_t       code_size;
    u8*          data;
    size_t       data_size;
} fl_vm;

fl_vm* fl_vm_create(size_t stack_size);
void   fl_vm_load_code(fl_vm* vm, fl_bytecode* code, size_t count);
void   fl_vm_load_data(fl_vm* vm, u8* data, size_t size);
int    fl_vm_exec(fl_vm* vm);
void   fl_vm_destroy(fl_vm* vm);

/******************************************************************************/

#endif /* vm.h */