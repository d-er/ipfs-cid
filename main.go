package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"ipfs-cid/internal"
	"runtime"
	"unsafe"
)

//export GetCid
func GetCid(name *C.char) *C.char {

	filename := C.GoString(name)

	cid := internal.Cid(filename)
	cCid := C.CString(cid)

	// to free the memory
	runtime.SetFinalizer(&cCid, func(ptr **C.char) {
		C.free(unsafe.Pointer(*ptr))
	})

	return cCid
}

func main() {}
