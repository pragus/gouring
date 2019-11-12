package main

import (
	"go_uring/uring"
	"log"
	"unsafe"
)

func main() {
	entries := uint(64)
	ring := &uring.Uring{}
	cqe := &uring.IoUringCqe{}
	ptr := unsafe.Pointer(&cqe)

	res := uring.QueueInit(ring, entries, 0)
	log.Printf("init, res: %+v, ring: %+v", res, ring)

	sqe := uring.SqeGet(ring)
	for i := uint(0); i < entries; i++ {
		uring.PrepOpNop(sqe)
	}

	uring.QueueSubmit(ring)
	res = uring.WaitCqe(ring, unsafe.Pointer(&ptr))
	log.Printf("wait_cqe, res: %+v", res)
	log.Printf("ring: %+v", ring)
	log.Printf("cqe: %+v", cqe)

	//ring.DoRing()
}
