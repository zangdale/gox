package xencode

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GB18030ToUTF8  GB18030 编码转 UTF8
func GB18030ToUTF8(output []byte) ([]byte, error) {
	return ioutil.ReadAll(transform.NewReader(bytes.NewBuffer(output), simplifiedchinese.GB18030.NewDecoder()))
}
