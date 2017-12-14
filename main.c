#include "exec/vm.h"
#include "internal/common.h"

int main(void)
{
    v_bytecode code[] = {
        Syscall,
        0,
    };

    fatalf("%zu\n", sizeof(v_bytecode));

    u8 data[] = {};

    // for (size_t i = 0; i < sizeof(data) / sizeof(data[0]); i++)
    //     printf("0x%02x ", data[i]);
    // printf("0x%x\n", *(u32*)(data));

    v_vm* vm = v_vm_create(512);
    v_vm_load_code(vm, code, sizeof(code));
    v_vm_load_data(vm, data, sizeof(data));
    v_vm_exec(vm);
    v_vm_destroy(vm);

    return 0;
}