package mem

import (
	"fmt"
	"runtime"
)

type Unit int

const (
	KiB Unit = 1024       // 1024 B
	MiB Unit = 1024 * KiB // 1024 KiB
)

func (u Unit) String() string {
	switch u {
	case KiB:
		return "KiB"
	case MiB:
		return "MiB"
	}

	return fmt.Sprintf("Unit(%d)", u)
}

// HeapMem gets bytes of allocated heap objects
func HeapMem() uint64 {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	return mem.HeapAlloc
}

// StackMem gets bytes of stack size
func StackMem() uint64 {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	return mem.StackSys
}

func FormatMem(unit Unit) string {
	return fmt.Sprintf("stack: %.2f %s, heap: %.2f %s", float32(StackMem())/float32(unit), unit, float32(HeapMem())/float32(unit), unit) //nolint
}
