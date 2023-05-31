/**
@Author: wei-G
@Email: wei_g_it@163.com
@Date: 2023/5/29 10:16
@Description:
*/

package buffer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCircularBuffer(t *testing.T) {
	bufferSize := 10
	cb := NewCircularBuffer(bufferSize)

	// Ensure that the circular buffer is initialized correctly
	if len(cb.buffer) != bufferSize || cb.start != 0 || cb.size != 0 {
		t.Errorf("NewCircularBuffer: incorrect initialization")
	}
}

func TestCircularBuffer_WriteRead(t *testing.T) {
	cb := NewCircularBuffer(5)
	data := []byte("Hello, World!")
	expected := []byte("orld!")

	cb.Write(data)
	result := cb.Read()

	// Ensure that the read data is correct
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Read: expected %v, got %v", expected, result)
	}
}

func TestCircularBuffer_Cap(t *testing.T) {
	bufferSize := 10
	cb := NewCircularBuffer(bufferSize)
	expected := bufferSize

	result := cb.Cap()

	// Ensure that the buffer size is correct
	if result != expected {
		t.Errorf("Size: expected %d, got %d", expected, result)
	}
}

func TestCircularBuffer_Reset(t *testing.T) {
	cb := NewCircularBuffer(5)
	data1 := []byte("Hello")
	data2 := []byte("Wor1")
	expected := []byte("Wor1")

	cb.Write(data1)
	cb.Reset()
	cb.Write(data2)
	result := cb.Read()

	fmt.Println(cb.String())
	// Ensure that only the latest data is read after reset and write
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reset: expected %v, got %v", expected, result)
	}
}

func TestCircularBuffer_WriteAndGetLen(t *testing.T) {
	cb := NewCircularBuffer(5)
	data := []byte("Hello")
	expectedLen := len(data)

	cb.Write(data)
	resultLen := cb.Len()

	if resultLen != expectedLen {
		t.Errorf("WriteAndGetLen: expected length %d, got %d", expectedLen, resultLen)
	}
}

func TestCircularBuffer_ResetAndWriteAndGetLen(t *testing.T) {
	cb := NewCircularBuffer(5)
	data1 := []byte("Hello")
	data2 := []byte("Wor1")
	expectedLen := len(data2)

	cb.Write(data1)
	cb.Reset()
	cb.Write(data2)
	resultLen := cb.Len()

	if resultLen != expectedLen {
		t.Errorf("ResetAndWriteAndGetLen: expected length %d, got %d", expectedLen, resultLen)
	}
}
