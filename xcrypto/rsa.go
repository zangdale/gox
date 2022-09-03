package xcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

//https://github.com/travist/jsencrypt

// 加密
func RSAEncrypt(origData, puKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(puKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse PKIX public key %s", err.Error())
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RSADecrypt(ciphertext, prkey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(prkey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//EncryptByRSA RSA加密 后 base64 再加
func EncryptByRSA(data []byte) (string, error) {
	res, err := RSAEncrypt(data, defaultPuKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

//DecryptByRSA RSA 解密
func DecryptByRSA(data string) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return RSADecrypt(dataByte, defaultPrKey)
}

var defaultPrKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCxS4S/GqCSzAedNFo2eXJ+DSw87DVJ8xCbqKD7hi2cr7kFD8B+
pjOKHjZc4f8JyEdZfX2zqoXQgSGDLtYA5lWLk6PQUBQhDYPB7nZwOsz1Wsv8AKM5
EbPTp3TcUGnRY4AF70SEoTRP3nqbaqIKuRq2FOHTd5sZjiv+gFLLXUH3VQIDAQAB
AoGAM0R6K2TAwBI9tWi5GX1+7RJUE33eXvbpe4mEm5cOQ3vQzbTjXfwjdTINWEiF
nkeK0kbmfXi23zcoAK4fdf0lCQ0y4vTDjPVDV36r9MAmdlHalA9vX4yNdGnjPK9q
nlzEbMY0RWG9ZeGPVq3kjQm1v/5SroMReQGwKPkIEVhgQtUCQQD/l9kof3qK6F1h
yQ6O1ahuqUnvaVWr1STWv06oeiS5LAKwNs1cohFO2oHzVaTQs0pmDyb3GnYJRIIF
Ty1sDZ2DAkEAsZPDuAixQIF5IuZCnT0ljebF3t+/AckAjPtwToF2IFCBtF/d5u7k
FN7X5HwVh1Iz570d0JhAKWwan0ixx70YRwJATJ96I5Dr7L6yWAFNUvascthfaN18
KHJSg+qAKzPK1JRkDe2v7QhNBgWtlYRkT4igUi5SsRuGrUqTbAILjOwb/wJBAI9v
xRL9anepXXjUN5CdGJ2Tf9c0Miw1+Qzn+OJg7lLR1MMnAK4N3wwAqLC1jgo9WxHg
D5ozsPgEi0iIRpoJYvcCQHGWzEdXGUvHTg2Rla1eDohQJagBW1/qSCHGpyAcMayq
AavvbzJqEABHZjtb7u1e7ms1erw2u1wvrqqR+FCSRcE=
-----END RSA PRIVATE KEY-----`)

var defaultPuKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCxS4S/GqCSzAedNFo2eXJ+DSw8
7DVJ8xCbqKD7hi2cr7kFD8B+pjOKHjZc4f8JyEdZfX2zqoXQgSGDLtYA5lWLk6PQ
UBQhDYPB7nZwOsz1Wsv8AKM5EbPTp3TcUGnRY4AF70SEoTRP3nqbaqIKuRq2FOHT
d5sZjiv+gFLLXUH3VQIDAQAB
-----END PUBLIC KEY-----`)

func SetKEY(prKey, puKey []byte) {
	defaultPrKey = prKey
	defaultPuKey = puKey
}
