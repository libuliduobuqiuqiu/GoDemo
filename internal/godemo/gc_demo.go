package godemo

import "runtime"

// UseAllocate keeps a small set of heap blocks live until they have been read.
// It demonstrates reachable allocations, not a predictable GC cycle count.
func UseAllocate() int {
	blocks := make([][]byte, 32)
	for i := range blocks {
		blocks[i] = make([]byte, 64<<10)
		blocks[i][0] = byte(i)
	}
	checksum := 0
	for _, block := range blocks {
		checksum += int(block[0])
	}
	runtime.KeepAlive(blocks)
	return checksum
}
