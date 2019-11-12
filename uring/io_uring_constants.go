package uring

// sqe->flags
const (
	IOSQE_FIXED_FILE = 0x1 /* use fixed fileset */
	IOSQE_IO_DRAIN   = 0x2 /* issue after inflight IO */
	IOSQE_IO_LINK    = 0x4 /* next IO depends on this one */
)

// io_uring_setup() flags
const (
	IORING_SETUP_IOPOLL = 0x1 /* io_context is polled */
	IORING_SETUP_SQPOLL = 0x2 /* SQ poll thread */
	IORING_SETUP_SQ_AFF = 0x4 /* sq_thread_cpu is valid */
)

// ops
const (
	IORING_OP_NOP             = 0x0
	IORING_OP_READV           = 0x1
	IORING_OP_WRITEV          = 0x2
	IORING_OP_FSYNC           = 0x3
	IORING_OP_READ_FIXED      = 0x4
	IORING_OP_WRITE_FIXED     = 0x5
	IORING_OP_POLL_ADD        = 0x6
	IORING_OP_POLL_REMOVE     = 0x7
	IORING_OP_SYNC_FILE_RANGE = 0x8
	IORING_OP_SENDMSG         = 0x9
	IORING_OP_RECVMSG         = 0xa
)

// sqe->fsync_flags
const (
	IORING_FSYNC_DATASYNC = 0x1
)

// Magic offsets for the application to mmap the data it needs
const (
	IORING_OFF_SQ_RING = 0x0
	IORING_OFF_CQ_RING = 0x8000000
	IORING_OFF_SQES    = 0x10000000
)

// sq_ring->flags
const (
	IORING_SQ_NEED_WAKEUP = 0x1 /* needs io_uring_enter wakeup */
)

// io_uring_enter(2) flags
const (
	IORING_ENTER_GETEVENTS = 0x1
	IORING_ENTER_SQ_WAKEUP = 0x2
)

// io_uring_register(2) opcodes and arguments
const (
	IORING_REGISTER_BUFFERS   = 0x0
	IORING_UNREGISTER_BUFFERS = 0x1
	IORING_REGISTER_FILES     = 0x2
	IORING_UNREGISTER_FILES   = 0x3
	IORING_REGISTER_EVENTFD   = 0x4
	IORING_UNREGISTER_EVENTFD = 0x5
)
