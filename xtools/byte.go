package xtools

// StringToBytes 非空字符串转为比特数组
func StringToBytes(s string) []byte {
	if s == "" {
		return nil
	}
	return []byte(s)
}
