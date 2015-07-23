package ozmecab

// #cgo CFLAGS: -I .
// #cgo LDFLAGS: -L /usr/local/mecab/lib -lmecab -lstdc++
// #include <mecab_c.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

import (
	"fmt"
	"strings"
)

func Init() error {
	v := C.Init()
	if v == 0 {
		return fmt.Errorf("C.Init()")
	}

	return nil
}

type Word struct {
	W string
	T string
	I string
}

func Parse(text string) ([]Word, error) {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	mw := make([]Word, 0)

	v := C.Parse(cs)
	if v == nil {
		return mw, fmt.Errorf("error")
	}

	gs := C.GoString(v)
	defer C.free(unsafe.Pointer(v))
	ss := strings.Split(gs, "\n")
	for _, i := range ss {
		sss := strings.Split(i, "\t")
		if len(sss) < 2 || len(sss[0]) == 0 {
			continue
		}

		var m Word
		m.W = sss[0]
		m.I = sss[1]
		ssss := strings.Split(sss[1], ",")
		if len(ssss) > 0 {
			m.T = ssss[0]
		}

		mw = append(mw, m)
	}

	return mw, nil
}

func Fin() {
	C.Fin()
}
