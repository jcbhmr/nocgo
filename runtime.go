//go:build !cgo

package nocgo

import (
	_ "runtime"
	_ "unsafe"
)

//go:linkname runtime_iscgo runtime.iscgo
var runtime_iscgo bool = true
var _ = runtime_iscgo
