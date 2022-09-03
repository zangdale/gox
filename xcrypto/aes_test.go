package xcrypto

import (
	"testing"
)

var data = []byte("ABCDABCDABCDABCD")

func TestAES(t *testing.T) {
	//加密
	str, _ := EncryptByAES(data, nil)
	//解密
	str1, _ := DecryptByAES(str, nil)
	//打印
	t.Logf(" 加密：%v\n 解密：%s\n ",
		str, str1,
	)
}

//测试速度
func BenchmarkAES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str, _ := EncryptByAES(data, nil)
		DecryptByAES(str, nil)
	}
}
