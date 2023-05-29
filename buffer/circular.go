/**
@Author: wei-G
@Email: wei_g_it@163.com
@Date: 2023/5/29 10:07
@Description: 支持并发访问的 环形缓冲区（Circular Buffer）实现
*/

package buffer

import (
	"sync"
)

type CircularBuffer struct {
	buffer []byte
	start  int
	size   int
	lock   sync.RWMutex
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		buffer: make([]byte, size),
		start:  0,
		size:   0,
	}
}

func (c *CircularBuffer) Write(data []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()

	dataSize := len(data)
	if dataSize > len(c.buffer) {
		// Data size is larger than the buffer size
		// Truncate the data to fit the buffer
		data = data[dataSize-len(c.buffer):]
		dataSize = len(data)
	}

	if c.size+dataSize <= len(c.buffer) {
		// Enough space available, write data directly
		end := (c.start + c.size) % len(c.buffer)
		copy(c.buffer[end:], data)
		c.size += dataSize
	} else {
		// Need to wrap around and overwrite old data
		remainingSize := len(c.buffer) - (c.size + dataSize)
		copy(c.buffer[c.start+c.size:], data[:remainingSize])
		copy(c.buffer, data[remainingSize:])
		c.start = (c.start + dataSize) % len(c.buffer)
		c.size = len(c.buffer)
	}
}

func (c *CircularBuffer) Read() []byte {
	c.lock.RLock()
	defer c.lock.RUnlock()

	data := make([]byte, c.size)
	if c.size > 0 {
		end := (c.start + c.size) % len(c.buffer)
		if c.start < end {
			copy(data, c.buffer[c.start:end])
		} else {
			copy(data, c.buffer[c.start:])
			copy(data[len(c.buffer)-c.start:], c.buffer[:end])
		}
	}
	return data
}

func (c *CircularBuffer) Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.buffer)
}

// Reset 重置缓冲区
func (c *CircularBuffer) Reset() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.start = 0
	c.size = 0
	// 没有重新 c.buffer 是为了 节省内存和提高效率
}

func (c *CircularBuffer) String() string {
	return string(c.Read())
}
