#include "vm.h"

/******************************************************************************/
fl_vm* v_vm_create(size_t stack_size)
{
    fl_vm* vm = malloc(sizeof(*vm));
    fl_malloc_check(vm, "fl_vm.create()");

    vm->stack = fl_stack_create(stack_size);
    return vm;
}

/******************************************************************************/
void v_vm_load_code(fl_vm* vm, fl_bytecode* code, size_t count)
{
    if (!vm)
        fatalf("fl_vm.load_code(%p, %p, size_t): null fl_vm reference\n", vm, code);

    if (!code)
        fatalf("fl_vm.load_code(%p, %p, size_t): no input bytecode\n", vm, code);

    vm->code_size = count;
    vm->code      = code;
}

/******************************************************************************/
void v_vm_load_data(fl_vm* vm, u8* data, size_t size)
{
    if (!vm)
        fatalf("fl_vm.load_data(%p, %p, size_t): null fl_vm reference\n", vm, data);

    if (!data) {
        vm->data      = NULL;
        vm->data_size = 0;
        return;
    }

    vm->data_size = size;
    vm->data      = data;
}

/******************************************************************************/
void v_vm_destroy(fl_vm* vm)
{
    if (!vm)
        return;

    fl_stack_destroy(vm->stack);
    vm->stack = NULL;
    fl_free(vm);
}

/******************************************************************************/
int v_vm_exec(fl_vm* vm)
{
    if (!vm)
        fatalf("fl_vm.exec(%p): null fl_vm reference\n", vm);

    if (!vm->code)
        fatalf("fl_vm.exec(%p): no input bytecode\n", vm);

    vm->pc = 0;

    return 0;
}