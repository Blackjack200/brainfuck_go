package bfruntime

type MemoryTape struct {
	tape map[uint64]byte
	pos  uint64
}

func (t *MemoryTape) lazy() {
	if t.tape == nil {
		t.tape = make(map[uint64]byte, 255)
	}
	if _, ok := t.tape[t.pos]; !ok {
		t.tape[t.pos] = 0
	}
}

func (t *MemoryTape) Next(delta uint64) {
	t.lazy()
	t.pos += delta
}

func (t *MemoryTape) Prev(delta uint64) {
	t.lazy()
	t.pos -= delta
}

func (t *MemoryTape) Write(b byte) {
	t.lazy()
	t.tape[t.pos] = b
}

func (t *MemoryTape) Read() byte {
	t.lazy()
	return t.tape[t.pos]
}

func (t *MemoryTape) Inc(delta uint64) {
	t.lazy()
	t.tape[t.pos] += byte(delta)
}

func (t *MemoryTape) Dec(delta uint64) {
	t.lazy()
	t.tape[t.pos] -= byte(delta)
}

func (t *MemoryTape) Dump() map[uint64]byte {
	n := make(map[uint64]byte)
	for k, v := range t.tape {
		n[k] = v
	}
	return n
}

func NewMemoryTape() *MemoryTape {
	return &MemoryTape{}
}
