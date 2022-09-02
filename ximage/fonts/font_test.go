package fonts

import (
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/zangdale/gox/ximage/fonts/cefcjk"
)

func TestGG(t *testing.T) {
	font, err := truetype.Parse(cefcjk.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	bg, err := gg.LoadImage("src.jpg")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(bg.Bounds().Max.X, bg.Bounds().Max.Y)

	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.DrawImage(bg, 0, 0)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored("Hello, 中文!", float64(bg.Bounds().Max.X/2),
		float64(bg.Bounds().Max.Y/2), 0.5, 0.5)
	dc.SavePNG("out.png")
}

func addLabel(img *image.RGBA, x, y int, label string) error {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	face, err := DefaultFace(48)
	if err != nil {
		return err
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}
	d.DrawString(label)
	return nil
}

func TestXimage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 600, 100))
	err := addLabel(img, 200, 30, "中文Hello")
	if err != nil {
		panic(err)
	}
	f, err := os.Create("hello-go.png")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
