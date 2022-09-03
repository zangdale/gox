package xcrypto

import (
	"fmt"
	"testing"
)

func Test_Rsa(t *testing.T) {
	etxt, err := EncryptByRSA([]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(etxt)

	decrypt, err := DecryptByRSA(etxt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(decrypt))
}
