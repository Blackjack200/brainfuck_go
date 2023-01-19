package bfruntime

import (
	"brainfuck/mem"
	"unsafe"
)

type CTape struct {
	idx uint64
	ptr unsafe.Pointer
}

func (t *CTape) Next(u uint64) {
	t.idx += u
}

func (t *CTape) Prev(u uint64) {
	t.idx -= u
}

func (t *CTape) Write(b byte) {
	mem.Set(t.ptr, t.idx, b)
}

func (t *CTape) Read() byte {
	return mem.Get(t.ptr, t.idx)
}

func (t *CTape) Inc(u uint64) {
	mem.Set(t.ptr, t.idx, mem.Get(t.ptr, t.idx)+byte(u))
}

func (t *CTape) Dec(u uint64) {
	mem.Set(t.ptr, t.idx, mem.Get(t.ptr, t.idx)-byte(u))
}

func (t *CTape) Dump() map[uint64]byte {
	m := make(map[uint64]byte)
	for i := uint64(0); i < 16; i++ {
		m[i] = mem.Get(t.ptr, i)
	}
	return m
}

func NewCTape() *CTape {
	c := mem.Create(4096)
	return &CTape{
		idx: 128,
		ptr: c,
	}
}
