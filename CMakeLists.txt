cmake_minimum_required(VERSION 3.9)
project(virt)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_COMPILER clang-5.0)
set(CMAKE_C_FLAGS "-Wall -Wextra -Werror")

add_executable(vexec
        internal/bytecode.h
        internal/common.h
        internal/memory.h
        internal/types.h
        internal/stack/stack.c
        internal/stack/stack.h
        exec/vm.c
        exec/vm.h
        main.c)