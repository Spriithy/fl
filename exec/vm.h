#ifndef EXEC_VM_H
#define EXEC_VM_H

#include "../internal/bytecode.h"
#include "../internal/common.h"
#include "../internal/stack/stack.h"

typedef struct {
    v_stack* stack;
    size_t pc;
    v_bytecode* code;
    size_t code_size;
    u8* data;
    size_t data_size;
} v_vm;

v_vm* v_vm_create(size_t stack_size);
void v_vm_load_code(v_vm* vm, v_bytecode* code, size_t count);
void v_vm_load_data(v_vm* vm, u8* data, size_t size);
int v_vm_exec(v_vm* vm);
void v_vm_destroy(v_vm* vm);

#endif /* vm.h */