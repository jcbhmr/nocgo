//go:build generate

package fakecgo

import "unsafe"

/*

*/
//go:generate go tool precgo "libc.so.6" -buildtags "!cgo && (darwin || freebsd || linux || netbsd)"
var _ [0]byte

type _ interface {
	malloc(size uintptr) unsafe.Pointer
	free(ptr unsafe.Pointer)
	setenv(name, value *byte, overwrite int32) int32
	unsetenv(name *byte) int32
	sigfillset(set *sigset_t) int32
	nanosleep(duration, rem *timespec) int32
	abort()
}

//go:generate go tool cgo-import-dynamic-bindgen "libpthread.so.0" -buildtags "!cgo && (darwin || freebsd || linux || netbsd)"
type _ interface {
	pthread_attr_init(attr *pthread_attr_t) int32
	pthread_create(thread *pthread_t, attr *pthread_attr_t, start_routine *byte, arg unsafe.Pointer) int32
	pthread_detach(thread *pthread_t) int32
	pthread_sigmask(how int32, set, oldset *sigset_t) int32
	pthread_self() pthread_t
	pthread_get_stacksize_np(thread *pthread_t) uintptr
	pthread_attr_getstacksize(attr *pthread_attr_t, size *uintptr) int32
	pthread_attr_setstacksize(attr *pthread_attr_t, size uintptr) int32
	pthread_attr_destroy(attr *pthread_attr_t) int32
	pthread_mutex_lock(mutex *pthread_mutex_t) int32
	pthread_mutex_unlock(mutex *pthread_mutex_t) int32
	pthread_cond_broadcast(cond *pthread_cond_t) int32
	pthread_setspecific(key pthread_key_t, value unsafe.Pointer) int32
}
