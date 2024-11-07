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
func GetCid(fileData *C.uchar, length C.int) *C.char {

	data := C.GoBytes(unsafe.Pointer(fileData), length)
	cid := internal.Cid(data)
	cCid := C.CString(cid)

	// to free the memory
	runtime.SetFinalizer(&cCid, func(ptr **C.char) {
		C.free(unsafe.Pointer(*ptr))
	})

	return cCid
}

func main() {}
