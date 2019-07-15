package utils

import (
	"encoding/binary"
)

// ConvInt64ToBytes convert int64 to []byte
func ConvInt64ToBytes(num int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(num))
	return b
}

// ConvBytesToInt64 convert []byte to int64
func ConvBytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
