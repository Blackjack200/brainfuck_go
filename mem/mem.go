package mem

/*
#include <stdlib.h>
#include <string.h>

typedef unsigned char CELL;

CELL* mem_arena_create(size_t size) {
    void* mem = malloc(size);
    if (mem == NULL) {
    return NULL;
    }
    memset(mem,0,sizeof(mem));
    return mem;
}

void mem_arena_destroy(CELL* arena) {
    free(arena);
}

CELL* mem_arena_get_ptr(unsigned char* arena, int offset) {
    return arena+offset;
}
*/
import "C"

import (
	"unsafe"
)

const byteSize = unsafe.Sizeof(byte(0))

func Create(i uint) unsafe.Pointer {
	return unsafe.Pointer(C.mem_arena_create(C.size_t(byteSize * uintptr(i))))
}

func Destroy(m unsafe.Pointer) {
	C.mem_arena_destroy((*C.uchar)(m))
}

func Get(m unsafe.Pointer, offset uint64) byte {
	return byte(*Addr(m, offset))
}

func Addr(m unsafe.Pointer, offset uint64) *C.uchar {
	return (*C.uchar)(C.mem_arena_get_ptr((*C.uchar)(m), C.int(offset)))
}

func Set(m unsafe.Pointer, offset uint64, val byte) {
	*Addr(m, offset) = C.uchar(val)
}
