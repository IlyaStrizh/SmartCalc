// go build -buildmode=plugin -o calc.so

package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L.
// #include "calc.h"
import "C"

import (
	"errors"
	"regexp"
	"strings"
	"unsafe"
)

func Calc(s *string, x string) (float64, error) {
	err := parser(s, x)
	if err == nil {
		var res C.double
		str := C.CString(*s)
		defer C.free(unsafe.Pointer(str))

		if err := C.calc(str, &res); err == 0 {
			return float64(res), nil
		} else {
			return 0.0, errors.New("error")
		}
	}

	return 0.0, err
}

func parser(s *string, x string) error {
	*s = strings.ReplaceAll(*s, "sqrt", "q")
	*s = strings.ReplaceAll(*s, "acos", "C")
	*s = strings.ReplaceAll(*s, "asin", "S")
	*s = strings.ReplaceAll(*s, "atan", "T")
	*s = strings.ReplaceAll(*s, "cos", "c")
	*s = strings.ReplaceAll(*s, "sin", "s")
	*s = strings.ReplaceAll(*s, "tan", "t")
	*s = strings.ReplaceAll(*s, "log", "L")
	*s = strings.ReplaceAll(*s, "ln", "l")
	*s = strings.ReplaceAll(*s, " ", "")
	*s = strings.ReplaceAll(*s, "x", x)
	match, _ := regexp.MatchString(
		`^[+\-*/^%cstCSTqLle.()0123456789]+$`, *s)
	if match {
		return nil
	}

	return errors.New("error")
}

func main() {}
