package uring

// Filled with the offset for mmap(2)
type SqOffsets struct {
	Head    uint32
	Tail    uint32
	Mask    uint32
	Entries uint32
	Flags   uint32
	Dropped uint32
	Array   uint32
	Resv1   uint32
	Resv2   uint64
}

// Filled with the offset for mmap(2)
type CqOffsets struct {
	Head     uint32
	Tail     uint32
	Mask     uint32
	Entries  uint32
	Overflow uint32
	Cqes     uint32
	Resv     [2]uint64
}

// Passed in for io_uring_setup(2). Copied back with updated info on success
type Params struct {
	SqEntries    uint32
	CqEntries    uint32
	Flags        uint32
	SqThreadCpu  uint32
	SqThreadIdle uint32
	Resv         [5]uint32
	SqOffsets    SqOffsets
	CqOffsets    CqOffsets
}
