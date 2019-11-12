package uring

/*
#cgo CFLAGS: -I. -I${SRCDIR}/liburing
#cgo LDFLAGS: -L. -L${SRCDIR}/liburing
#include <liburing.h>
*/
import "C"
import "unsafe"

type Cgo_IoUringSq C.struct_io_uring_sq
type Cgo_IoUringCq C.struct_io_uring_cq

type Cgo_IoUring C.struct_io_uring

func (r *Cgo_IoUring) toIoUring() *Uring {
	return (*Uring)(unsafe.Pointer(r))
}

type Cgo_IoUringSqe C.struct_io_uring_sqe
type Cgo_IoUringCqe C.struct_io_uring_cqe

type UringSq struct {
	Khead         *uint32
	Ktail         *uint32
	Kring_mask    *uint32
	Kring_entries *uint32
	Kflags        *uint32
	Kdropped      *uint32
	Array         *uint32
	Sqes          *IoUringSqe
	Sqe_head      uint32
	Sqe_tail      uint32
	Ring_sz       uint64
	Ring_ptr      *byte
}
type UringCq struct {
	Khead         *uint32
	Ktail         *uint32
	Kring_mask    *uint32
	Kring_entries *uint32
	Koverflow     *uint32
	Cqes          *IoUringCqe
	Ring_sz       uint64
	Ring_ptr      *byte
}
type Uring struct {
	Sq    UringSq
	Cq    UringCq
	Flags uint32
	Fd    int32
}

type IoUringSqe struct {
	Opcode    uint8
	Flags     uint8
	Ioprio    uint16
	Fd        int32
	Off       uint64
	Addr      uint64
	Len       uint32
	Rw_flags  int32
	User_data uint64
	Buf_index uint16
	Pad_cgo_0 [22]byte
}
type IoUringCqe struct {
	Data  uint64
	Res   int32
	Flags uint32
}
