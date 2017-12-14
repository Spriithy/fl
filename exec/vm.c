#include "vm.h"

/******************************************************************************/
v_vm* v_vm_create(size_t stack_size)
{
    v_vm* vm = malloc(sizeof(*vm));
    v_malloc_check(vm, "v_vm.create()");

    vm->stack = v_stack_create(stack_size);
    return vm;
}

/******************************************************************************/
void v_vm_load_code(v_vm* vm, v_bytecode* code, size_t count)
{
    if (!vm)
        fatalf("v_vm.load_code(%p, %p, size_t): null v_vm reference\n", vm, code);

    if (!code)
        fatalf("v_vm.load_code(%p, %p, size_t): no input bytecode\n", vm, code);

    vm->code_size = count;
    vm->code = code;
}

/******************************************************************************/
void v_vm_load_data(v_vm* vm, u8* data, size_t size)
{
    if (!vm)
        fatalf("v_vm.load_data(%p, %p, size_t): null v_vm reference\n", vm, data);

    if (!data) {
        vm->data = NULL;
        vm->data_size = 0;
        return;
    }

    vm->data_size = size;
    vm->data = data;
}

/******************************************************************************/
void v_vm_destroy(v_vm* vm)
{
    if (!vm)
        return;

    v_stack_destroy(vm->stack);
    vm->stack = NULL;
    v_free(vm);
}

/******************************************************************************/
int v_vm_exec(v_vm* vm)
{
    if (!vm)
        fatalf("v_vm.exec(%p): null v_vm reference\n", vm);

    if (!vm->code)
        fatalf("v_vm.exec(%p): no input bytecode\n", vm);

    vm->pc = 0;

    return 0;
}