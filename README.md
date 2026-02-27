# nocgo

## Installation

```sh
go get go.jcbhmr.com/nocgo
```

## Usage

- **Go's `import "C"` pseudo-package:** Post-processed by Go's `cgo` command to generate Go and C code.
- **Go's `cgo` command:** Generates Go and C code that work together using `//go:cgo_*` directives, the `runtime/cgo` package internals, and generated wrappers.
- **Go's `runtime/cgo` package:** Provides a `cgo.Handle` type that can round-trip Go values to C and back. Internally provides a bunch of C-specific Go interop functionality that ties into the Go runtime.
- **nocgo's `nocgo` command:** Generates Go code that uses `//go:cgo_import_dynamic` and `//go:linkname` to get raw C function pointers to shared object functions, the `nocgo` package internals, and generated wrappers.
- **nocgo's root package:** Mirrors the `runtime/cgo` package's `cgo.Handle` type and provides `Internal*` functions/types that are used by the generated code from the `nocgo` command.
- **nocgo's `xnocgo` package:** Similar to `go.jcbhmr.com/xstd/xruntime/xcgo` but for `CGO_ENABLED=0` environments, provides `xnocgo.Call` to call C function pointers and `xnocgo.NewCallback` to create C function pointers that call Go functions.

```go
import "go.jcbhmr.com/nocgo/c"
cString := c.CString("Hello!")
goString := c.GoString(cString)
```

```go
//go:build generate

/*

*/
//go:generate go run go.jcbhmr.com/nocgo/cmd/nocgo -buildtags "linux" -sofile "libc.so.6"
//go:generate go run go.jcbhmr.com/nocgo/cmd/nocgo -buildtags "darwin" -sofile "libSystem.B.dylib"
var _ = 0
```

```go
//go:build linux

//go:cgo_import_dynamic _libc_so_6__free free "libc.so.6"
//go:linkname _libc_so_6__free _libc_so_6__free
var _libc_so_6__free byte
var _libc_so_6__freeABI0 = uintptr(unsafe.Pointer(&_libc_so_6__free))

func free(ptr unsafe.Pointer) {
    
}
```

## Development

- **nocgo:** Should cover everything that [runtime/cgo](https://golang.org/pkg/runtime/cgo) does.