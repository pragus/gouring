package uring

import (
	"log"
	"unsafe"
)

func QueueSetup(entries uint, params *Params) int {
	return cgo_io_uring_setup(entries, params)
}

func QueueMmap(ring *Uring, fd int, params *Params) int {
	return cgo_io_uring_queue_mmap(ring, fd, params)
}

func QueueInit(ring *Uring, entries, flags uint) int {
	return cgo_io_uring_queue_init(ring, entries, flags)
}

func QueueExit(ring *Uring) {
	cgo_io_uring_queue_exit(ring)
}

func QueueSubmit(ring *Uring) int {
	return cgo_io_uring_submit(ring)

}

func WaitCqe(r *Uring, p unsafe.Pointer) int {
	return cgo_io_uring_wait_cqe(r, p)
}

func SqeGet(ring *Uring) *IoUringSqe {
	return cgo_io_uring_get_sqe(ring)

}

func PrepOpNop(sqe *IoUringSqe) {
	cgo_io_uring_prep_nop(sqe)
}

func DoRing() {
	ring := &Uring{}
	res := QueueInit(ring, 32, 0)
	log.Printf("init, res: %+v, ring: %+v", res, ring)
}
