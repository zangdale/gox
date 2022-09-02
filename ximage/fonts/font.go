package fonts

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/zangdale/gox/ximage/fonts/cefcjk"
)

var DefaultFontData = cefcjk.TTF

// DefaultFont 加载默认的字体
func DefaultFont() (font *truetype.Font, err error) {
	return truetype.Parse(DefaultFontData)
}

// DefaultFace 返回默认字体对象
func DefaultFace(size float64) (font.Face, error) {
	f, err := DefaultFont()
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(f, &truetype.Options{Size: size}), nil
}
