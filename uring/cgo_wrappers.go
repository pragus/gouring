package uring

/*
#cgo CFLAGS: -I. -I${SRCDIR}/c/liburing/src/include -static
#cgo LDFLAGS: -L. -L${SRCDIR}/c/liburing/src

#include <linux/fs.h>
#include <linux/types.h>
#include <liburing.h>

extern int io_uring_setup(unsigned entries, struct io_uring_params *p);
extern int io_uring_enter(unsigned fd, unsigned to_submit,
	unsigned min_complete, unsigned flags, sigset_t *sig);
extern int io_uring_register(int fd, unsigned int opcode, const void *arg,
	unsigned int nr_args);

extern int io_uring_queue_init_params(unsigned entries, struct io_uring *ring,
	struct io_uring_params *p);
extern int io_uring_queue_init(unsigned entries, struct io_uring *ring,
	unsigned flags);
extern int io_uring_queue_mmap(int fd, struct io_uring_params *p,
	struct io_uring *ring);
extern void io_uring_queue_exit(struct io_uring *ring);
unsigned io_uring_peek_batch_cqe(struct io_uring *ring,
	struct io_uring_cqe **cqes, unsigned count);
extern int io_uring_wait_cqes(struct io_uring *ring,
	struct io_uring_cqe **cqe_ptr, unsigned wait_nr,
	struct __kernel_timespec *ts, sigset_t *sigmask);
extern int io_uring_wait_cqe_timeout(struct io_uring *ring,
	struct io_uring_cqe **cqe_ptr, struct __kernel_timespec *ts);
extern int io_uring_submit(struct io_uring *ring);
extern int io_uring_submit_and_wait(struct io_uring *ring, unsigned wait_nr);
extern struct io_uring_sqe *io_uring_get_sqe(struct io_uring *ring);

extern int io_uring_register_buffers(struct io_uring *ring,
					const struct iovec *iovecs,
					unsigned nr_iovecs);
extern int io_uring_unregister_buffers(struct io_uring *ring);
extern int io_uring_register_files(struct io_uring *ring, const int *files,
					unsigned nr_files);
extern int io_uring_unregister_files(struct io_uring *ring);
extern int io_uring_register_files_update(struct io_uring *ring, unsigned off,
					int *files, unsigned nr_files);
extern int io_uring_register_eventfd(struct io_uring *ring, int fd);
extern int io_uring_unregister_eventfd(struct io_uring *ring);

*/
import "C"
import (
	"unsafe"
)

/*
 * sqe->flags
 */
const (
	Cgo_IOSQE_FIXED_FILE = C.IOSQE_FIXED_FILE /* use fixed fileset */
	Cgo_IOSQE_IO_DRAIN   = C.IOSQE_IO_DRAIN   /* issue after inflight IO */
	Cgo_IOSQE_IO_LINK    = C.IOSQE_IO_LINK    /* next IO depends on this one */
)

/*
 * io_uring_setup() flags
 */

const (
	Cgo_IORING_SETUP_IOPOLL = C.IORING_SETUP_IOPOLL /* io_context is polled */
	Cgo_IORING_SETUP_SQPOLL = C.IORING_SETUP_SQPOLL /* SQ poll thread */
	Cgo_IORING_SETUP_SQ_AFF = C.IORING_SETUP_SQ_AFF /* sq_thread_cpu is valid */
	Cgo_IORING_SETUP_CQSIZE = C.IORING_SETUP_CQSIZE /* app defines CQ size */
)

const (
	Cgo_IORING_OP_NOP             = C.IORING_OP_NOP
	Cgo_IORING_OP_READV           = C.IORING_OP_READV
	Cgo_IORING_OP_WRITEV          = C.IORING_OP_WRITEV
	Cgo_IORING_OP_FSYNC           = C.IORING_OP_FSYNC
	Cgo_IORING_OP_READ_FIXED      = C.IORING_OP_READ_FIXED
	Cgo_IORING_OP_WRITE_FIXED     = C.IORING_OP_WRITE_FIXED
	Cgo_IORING_OP_POLL_ADD        = C.IORING_OP_POLL_ADD
	Cgo_IORING_OP_POLL_REMOVE     = C.IORING_OP_POLL_REMOVE
	Cgo_IORING_OP_SYNC_FILE_RANGE = C.IORING_OP_SYNC_FILE_RANGE
	Cgo_IORING_OP_SENDMSG         = C.IORING_OP_SENDMSG
	Cgo_IORING_OP_RECVMSG         = C.IORING_OP_RECVMSG
	Cgo_IORING_OP_TIMEOUT         = C.IORING_OP_TIMEOUT
	Cgo_IORING_OP_TIMEOUT_REMOVE  = C.IORING_OP_TIMEOUT_REMOVE
	Cgo_IORING_OP_ACCEPT          = C.IORING_OP_ACCEPT
	Cgo_IORING_OP_ASYNC_CANCEL    = C.IORING_OP_ASYNC_CANCEL
	Cgo_IORING_OP_LINK_TIMEOUT    = C.IORING_OP_LINK_TIMEOUT
)

/*
 * sqe->fsync_flags
 */
const (
	Cgo_IORING_FSYNC_DATASYNC = C.IORING_FSYNC_DATASYNC
)

/*
 * sqe->timeout_flags
 */
const (
	Cgo_IORING_TIMEOUT_ABS = C.IORING_TIMEOUT_ABS
)

/*
 * Magic offsets for the application to mmap the data it needs
 */

const (
	Cgo_IORING_OFF_SQ_RING = C.IORING_OFF_SQ_RING
	Cgo_IORING_OFF_CQ_RING = C.IORING_OFF_CQ_RING
	Cgo_IORING_OFF_SQES    = C.IORING_OFF_SQES
)

/*
 * Filled with the offset for mmap(2)
 */

type Cgo_io_uring_sq C.struct_io_uring_sq
type Cgo_io_uring_cq C.struct_io_uring_cq

type Cgo_io_uring_sqe C.struct_io_uring_sqe
type Cgo_io_uring_cqe C.struct_io_uring_cqe

type Cgo_io_sqring_offsets C.struct_io_sqring_offsets
type Cgo_io_cqring_offsets C.struct_io_cqring_offsets

/*
 * Passed in for io_uring_setup(2). Copied back with updated info on success
 */
type Cgo_io_uring_params C.struct_io_uring_params

/*
 * sq_ring->flags
 */
const (
	Cgo_IORING_SQ_NEED_WAKEUP = C.IORING_SQ_NEED_WAKEUP /* needs io_uring_enter wakeup */
)

/*
 * io_uring_enter(2) flags
 */
const (
	Cgo_IORING_ENTER_GETEVENTS = C.IORING_ENTER_GETEVENTS
	Cgo_IORING_ENTER_SQ_WAKEUP = C.IORING_ENTER_SQ_WAKEUP
)

/*
 * io_uring_register(2) opcodes and arguments
 */
const (
	Cgo_IORING_REGISTER_BUFFERS      = C.IORING_REGISTER_BUFFERS
	Cgo_IORING_UNREGISTER_BUFFERS    = C.IORING_UNREGISTER_BUFFERS
	Cgo_IORING_REGISTER_FILES        = C.IORING_REGISTER_FILES
	Cgo_IORING_UNREGISTER_FILES      = C.IORING_UNREGISTER_FILES
	Cgo_IORING_REGISTER_EVENTFD      = C.IORING_REGISTER_EVENTFD
	Cgo_IORING_UNREGISTER_EVENTFD    = C.IORING_UNREGISTER_EVENTFD
	Cgo_IORING_REGISTER_FILES_UPDATE = C.IORING_REGISTER_FILES_UPDATE
)

type Cgo_io_uring_files_update C.struct_io_uring_files_update

// DEF: extern int io_uring_setup(unsigned entries, struct Params *p);
func cgo_io_uring_setup(entries uint, params *Params) int {
	cgoEntries := C.uint(entries)
	cgoParams := (*C.struct_io_uring_params)(unsafe.Pointer(params))
	res := C.io_uring_setup(cgoEntries, cgoParams)
	return int(res)
}

// DEF: extern int io_uring_queue_mmap(int fd, struct Params *p, struct io_uring *ring);
func cgo_io_uring_queue_mmap(r *Uring, fd int, params *Params) int {
	cgoFd := C.int(fd)
	cgoParams := (*C.struct_io_uring_params)(unsafe.Pointer(params))
	cgoRing := (*C.struct_io_uring)(unsafe.Pointer(r))
	res := C.io_uring_queue_mmap(cgoFd, cgoParams, cgoRing)
	return int(res)

}

// DEF: extern int io_uring_queue_init(unsigned entries, struct io_uring *ring, unsigned flags);
func cgo_io_uring_queue_init(r *Uring, entries, flags uint) int {
	cgoRing := (*C.struct_io_uring)(unsafe.Pointer(r))
	cgoEntries := C.uint(entries)
	cgoFlags := C.uint(flags)
	res := C.io_uring_queue_init(cgoEntries, cgoRing, cgoFlags)
	return int(res)
}

// DEF: extern void io_uring_queue_exit(struct io_uring *ring);
func cgo_io_uring_queue_exit(r *Uring) {
	cgoRing := (*C.struct_io_uring)(unsafe.Pointer(r))
	C.io_uring_queue_exit(cgoRing)
}

// DEF: extern int io_uring_submit(struct io_uring *ring);
func cgo_io_uring_submit(r *Uring) int {
	cgoRing := (*C.struct_io_uring)(unsafe.Pointer(r))
	res := C.io_uring_submit(cgoRing)
	return int(res)
}

// DEF: extern int io_uring_wait_cqe(struct io_uring *ring, struct io_uring_cqe **cqe_ptr);
func cgo_io_uring_wait_cqe(r *Uring, p unsafe.Pointer) int {
	cgoRing := (*C.struct_io_uring)(unsafe.Pointer(r))
	cgoCqePtr := (**C.struct_io_uring_cqe)(p)
	res := C.io_uring_wait_cqe(cgoRing, cgoCqePtr)
	return int(res)
}

// DEF: static inline void io_uring_prep_nop(struct io_uring_sqe *sqe)
func cgo_io_uring_prep_nop(sqe *IoUringSqe) {
	cgoSqe := (*C.struct_io_uring_sqe)(unsafe.Pointer(sqe))
	C.io_uring_prep_nop(cgoSqe)
}

func cgo_io_uring_get_sqe(r *Uring) *IoUringSqe {
	cgoRing := (*C.struct_io_uring)(unsafe.Pointer(r))
	cgoSqe := unsafe.Pointer(C.io_uring_get_sqe(cgoRing))
	return (*IoUringSqe)(cgoSqe)
}
