/**
@Author: wei-G
@Email: wei_g_it@163.com
@Date: 2023/4/6 10:18
@Description:
*/

package byteconv

// GiBToByte 将 GiB 转换为 字节
func GiBToByte(input int64) int64 {
	return input * 1024 * 1024 * 1024
}

// ByteToGiB 将 字节 转换为 GiB
func ByteToGiB(input int64) float64 {
	return float64(input) / 1024 / 1024 / 1024
}

// MiBToByte 将 MiB 转换为 字节
func MiBToByte(input int64) int64 {
	return input * 1024 * 1024
}

// ByteToMiB 将 字节 转换为 MiB
func ByteToMiB(input int64) float64 {
	return float64(input) / 1024 / 1024
}

// KiBToByte 将 KiB 转换为 字节
func KiBToByte(input int64) int64 {
	return input * 1024
}

// ByteToKiB 将 字节 转换为 KiB
func ByteToKiB(input int64) float64 {
	return float64(input) / 1024
}

// GBToByte 将 GB 转换为 字节
func GBToByte(input int64) int64 {
	return input * 1000 * 1000 * 1000
}

// ByteToGB 将 字节 转换为 GB
func ByteToGB(input int64) float64 {
	return float64(input) / 1000 / 1000 / 1000
}

// MBToByte 将 MB 转换为 字节
func MBToByte(input int64) int64 {
	return input * 1000 * 1000
}

// ByteToMB 将 字节 转换为 MB
func ByteToMB(input int64) float64 {
	return float64(input) / 1000 / 1000
}

// KBToByte 将 KB 转换为 字节
func KBToByte(input int64) int64 {
	return input * 1024
}

// ByteToKB 将 字节 转换为 KB
func ByteToKB(input int64) float64 {
	return float64(input) / 1024
}
