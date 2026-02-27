//go:build !cgo

package nocgo

import (
	_ "syscall"
	"unsafe"
)

//go:cgo_import_dynamic _libc_setegid_LOCATION setegid "libc.so.6"
//go:linkname _libc_setegid_LOCATION _libc_setegid_LOCATION
//go:linkname syscall_cgo_libc_setegid syscall.cgo_libc_setegid
var _libc_setegid_LOCATION byte
var syscall_cgo_libc_setegid = unsafe.Pointer(&_libc_setegid_LOCATION)
var _ = syscall_cgo_libc_setegid

//go:cgo_import_dynamic _libc_seteuid_LOCATION seteuid "libc.so.6"
//go:linkname _libc_seteuid_LOCATION _libc_seteuid_LOCATION
//go:linkname syscall_cgo_libc_seteuid syscall.cgo_libc_seteuid
var _libc_seteuid_LOCATION byte
var syscall_cgo_libc_seteuid = unsafe.Pointer(&_libc_seteuid_LOCATION)
var _ = syscall_cgo_libc_seteuid

//go:cgo_import_dynamic _libc_setregid_LOCATION setregid "libc.so.6"
//go:linkname _libc_setregid_LOCATION _libc_setregid_LOCATION
//go:linkname syscall_cgo_libc_setregid syscall.cgo_libc_setregid
var _libc_setregid_LOCATION byte
var syscall_cgo_libc_setregid = unsafe.Pointer(&_libc_setregid_LOCATION)
var _ = syscall_cgo_libc_setregid

//go:cgo_import_dynamic _libc_setresgid_LOCATION setresgid "libc.so.6"
//go:linkname _libc_setresgid_LOCATION _libc_setresgid_LOCATION
//go:linkname syscall_cgo_libc_setresgid syscall.cgo_libc_setresgid
var _libc_setresgid_LOCATION byte
var syscall_cgo_libc_setresgid = unsafe.Pointer(&_libc_setresgid_LOCATION)
var _ = syscall_cgo_libc_setresgid

//go:cgo_import_dynamic _libc_setresuid_LOCATION setresuid "libc.so.6"
//go:linkname _libc_setresuid_LOCATION _libc_setresuid_LOCATION
//go:linkname syscall_cgo_libc_setresuid syscall.cgo_libc_setresuid
var _libc_setresuid_LOCATION byte
var syscall_cgo_libc_setresuid = unsafe.Pointer(&_libc_setresuid_LOCATION)
var _ = syscall_cgo_libc_setresuid

//go:cgo_import_dynamic _libc_setreuid_LOCATION setreuid "libc.so.6"
//go:linkname _libc_setreuid_LOCATION _libc_setreuid_LOCATION
//go:linkname syscall_cgo_libc_setreuid syscall.cgo_libc_setreuid
var _libc_setreuid_LOCATION byte
var syscall_cgo_libc_setreuid = unsafe.Pointer(&_libc_setreuid_LOCATION)
var _ = syscall_cgo_libc_setreuid

//go:cgo_import_dynamic _libc_setgroups_LOCATION setgroups "libc.so.6"
//go:linkname _libc_setgroups_LOCATION _libc_setgroups_LOCATION
//go:linkname syscall_cgo_libc_setgroups syscall.cgo_libc_setgroups
var _libc_setgroups_LOCATION byte
var syscall_cgo_libc_setgroups = unsafe.Pointer(&_libc_setgroups_LOCATION)
var _ = syscall_cgo_libc_setgroups

//go:cgo_import_dynamic _libc_setgid_LOCATION setgid "libc.so.6"
//go:linkname _libc_setgid_LOCATION _libc_setgid_LOCATION
//go:linkname syscall_cgo_libc_setgid syscall.cgo_libc_setgid
var _libc_setgid_LOCATION byte
var syscall_cgo_libc_setgid = unsafe.Pointer(&_libc_setgid_LOCATION)
var _ = syscall_cgo_libc_setgid

//go:cgo_import_dynamic _libc_setuid_LOCATION setuid "libc.so.6"
//go:linkname _libc_setuid_LOCATION _libc_setuid_LOCATION
//go:linkname syscall_cgo_libc_setuid syscall.cgo_libc_setuid
var _libc_setuid_LOCATION byte
var syscall_cgo_libc_setuid = unsafe.Pointer(&_libc_setuid_LOCATION)
var _ = syscall_cgo_libc_setuid
