package cryptonight

// #cgo CFLAGS: -maes
// #include "cryptonight.h"
import "C"
import "unsafe"

func Hash(d []byte) [32]byte {
	return doHash(d, true)
}

func FastHash(d []byte) [32]byte {
	return doHash(d, false)
}

func doHash(d []byte, slow bool) [32]byte {
	l := len(d)
	b := make([]C.char, l)
	for i, c := range d {
		b[i] = C.char(c)
	}
	var cr [32]C.char
	bptr := unsafe.Pointer(&b[0])
	if slow {
		C.cn_slow_hash(bptr, C.size_t(l), &cr[0])
	} else {
		C.cn_fast_hash(bptr, C.size_t(l), &cr[0])
	}
	var r [32]byte
	for i, c := range cr {
		r[i] = byte(c)
	}
	return r
}
