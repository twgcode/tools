/**
@Author: wei-G
@Email: wei_g_it@163.com
@Date: 2023/4/6 10:18
@Description:
*/

package byteconv

// GBToByte 将GB转换为字节
func GBToByte(gb int64) int64 {
	return gb * 1000 * 1000 * 1000
}

// ByteToGB 将字节转换为GB
func ByteToGB(input int64) float64 {
	return float64(input) / 1000 / 1000 / 1000
}
