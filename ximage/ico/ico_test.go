package ico

import (
	"bytes"
	"image"
	"io/ioutil"
	"testing"
)

func TestICO(t *testing.T) {
	rf, err := ioutil.ReadFile("./test.ico")
	if err != nil {
		t.Fatal(err)
	}
	no, err := imageCompress_no(rf)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("./_test.ico", no, 0644)
	if err != nil {
		t.Fatal(err)
	}
}

// 无效
func imageCompress_no(dataI []byte) (dataO []byte, err error) {

	var origin image.Image
	origin, err = Decode(bytes.NewReader(dataI))
	if err != nil {
		return nil, err
	}

	// resize "github.com/nfnt/resize"
	// origin = resize.Thumbnail(32, 32, origin, resize.NearestNeighbor)

	var data = new(bytes.Buffer)
	err = Encode(data, origin)
	if err != nil {
		return nil, err
	}

	return data.Bytes(), err
}
