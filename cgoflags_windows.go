// +build windows

package python

// #cgo CFLAGS: -IC:/Python27/include -DMS_WIN64
// #cgo LDFLAGS: -LC:/Python27/libs -lpython27
// #include "go-python.h"
import "C"

// EOF
