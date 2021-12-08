package cgo

// void hello();
import "C"

func Hello() {
	C.hello()
}
