package ico

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"image"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"

	"golang.org/x/image/bmp"
)

func init() {
	image.RegisterFormat("ico", "\x00\x00\x01\x00", Decode, DecodeConfig)
}

type icondir struct {
	Reserved uint16
	Type     uint16
	Count    uint16
}

type icondirEntry struct {
	Width        byte
	Height       byte
	PaletteCount byte
	Reserved     byte

	ColorPlanes  uint16
	BitsPerPixel uint16
	Size         uint32
	Offset       uint32
}

func FindBestIcon(entries []*icondirEntry) *icondirEntry {
	if len(entries) == 0 {
		return nil
	}

	best := entries[0]
	for _, e := range entries {
		if (e.width() > best.width()) && (e.height() > best.height()) {
			best = e
		}
	}
	return best
}

// parseIco 解析图标并返回图标的元信息
func parseIco(r io.Reader) (*icondir, []*icondirEntry, error) {
	dir := &icondir{}
	var entries []*icondirEntry

	var err error
	err = binary.Read(r, binary.LittleEndian, &dir.Reserved)
	if err != nil {
		return nil, nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &dir.Type)
	if err != nil {
		return nil, nil, err
	}

	err = binary.Read(r, binary.LittleEndian, &dir.Count)
	if err != nil {
		return nil, nil, err
	}

	for i := uint16(0); i < dir.Count; i++ {
		entry := &icondirEntry{}
		e := parseIcondirEntry(r, entry)
		if e != nil {
			return nil, nil, e
		}
		entries = append(entries, entry)
	}

	return dir, entries, err
}

func parseIcondirEntry(r io.Reader, e *icondirEntry) error {
	err := binary.Read(r, binary.LittleEndian, e)
	if err != nil {
		return err
	}

	return nil
}

type dibHeader struct {
	dibHeaderSize uint32
	width         uint32
	height        uint32
}

func (e *icondirEntry) ColorCount() int {
	if e.PaletteCount == 0 {
		return 256
	}
	return int(e.PaletteCount)
}

func (e *icondirEntry) width() int {
	if e.Width == 0 {
		return 256
	}
	return int(e.Width)
}

func (e *icondirEntry) height() int {
	if e.Height == 0 {
		return 256
	}
	return int(e.Height)
}

// DecodeConfig 只返回最大图像的尺寸包含在图标中而不解码整个图标文件。
func DecodeConfig(r io.Reader) (image.Config, error) {
	_, entries, err := parseIco(r)
	if err != nil {
		return image.Config{}, err
	}

	best := FindBestIcon(entries)
	if best == nil {
		return image.Config{}, errInvalid
	}
	return image.Config{Width: best.width(), Height: best.height()}, nil
}

// 从icondirEntry读取的位图头结构
type bitmapHeaderRead struct {
	Size            uint32
	Width           uint32
	Height          uint32
	Planes          uint16
	BitCount        uint16
	Compression     uint32
	ImageSize       uint32
	XPixelsPerMeter uint32
	YPixelsPerMeter uint32
	ColorsUsed      uint32
	ColorsImportant uint32
}

// 生成的位图头结构  bmp.Decode()
type bitmapHeaderWrite struct {
	sigBM           [2]byte
	fileSize        uint32
	resverved       [2]uint16
	pixOffset       uint32
	Size            uint32
	Width           uint32
	Height          uint32
	Planes          uint16
	BitCount        uint16
	Compression     uint32
	ImageSize       uint32
	XPixelsPerMeter uint32
	YPixelsPerMeter uint32
	ColorsUsed      uint32
	ColorsImportant uint32
}

var errInvalid = errors.New("ico: invalid ICO image")

// Decode 返回图标中包含的最大图像，可能是BMP或PNG
func Decode(r io.Reader) (image.Image, error) {
	icoBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	r = bytes.NewReader(icoBytes)
	_, entries, err := parseIco(r)
	if err != nil {
		return nil, errInvalid
	}

	best := FindBestIcon(entries)
	if best == nil {
		return nil, errInvalid
	}

	return parseImage(best, icoBytes)
}

func parseImage(entry *icondirEntry, icoBytes []byte) (image.Image, error) {
	r := bytes.NewReader(icoBytes)
	r.Seek(int64(entry.Offset), 0)

	// 先尝试PNG，然后再尝试BMP
	img, err := png.Decode(r)
	if err != nil {
		return parseBMP(entry, icoBytes)
	}
	return img, nil
}

func parseBMP(entry *icondirEntry, icoBytes []byte) (image.Image, error) {
	bmpBytes, err := makeFullBMPBytes(entry, icoBytes)
	if err != nil {
		return nil, err
	}
	return bmp.Decode(bmpBytes)
}

func makeFullBMPBytes(entry *icondirEntry, icoBytes []byte) (*bytes.Buffer, error) {
	r := bytes.NewReader(icoBytes)
	r.Seek(int64(entry.Offset), 0)

	var err error
	h := bitmapHeaderRead{}

	err = binary.Read(r, binary.LittleEndian, &h)
	if err != nil {
		return nil, err
	}

	if h.Size != 40 || h.Planes != 1 {
		return nil, errInvalid
	}

	var pixOffset uint32
	if h.ColorsUsed == 0 && h.BitCount <= 8 {
		pixOffset = 14 + 40 + 4*(1<<h.BitCount)
	} else {
		pixOffset = 14 + 40 + 4*h.ColorsUsed
	}

	writeHeader := &bitmapHeaderWrite{
		sigBM:           [2]byte{'B', 'M'},
		fileSize:        14 + 40 + uint32(len(icoBytes)), // correct? important?
		pixOffset:       pixOffset,
		Size:            40,
		Width:           uint32(h.Width),
		Height:          uint32(h.Height / 2),
		Planes:          h.Planes,
		BitCount:        h.BitCount,
		Compression:     h.Compression,
		ColorsUsed:      h.ColorsUsed,
		ColorsImportant: h.ColorsImportant,
	}

	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, writeHeader); err != nil {
		return nil, err
	}
	io.CopyN(buf, r, int64(entry.Size))

	return buf, nil
}

func newIcondir() icondir {
	var id icondir
	id.Type = 1
	id.Count = 1
	return id
}

func newIcondirentry() icondirEntry {
	var ide icondirEntry
	ide.ColorPlanes = 1   // Windows应该不介意0或1，但其他图标文件似乎在这里有1
	ide.BitsPerPixel = 32 // 位图为24,png为24/32。 现在设置为32
	ide.Offset = 22       //6 icondir + 16 icondirentry，下一个图像将是这个图像的大小 size + 16 icondirentry，以此类推
	return ide
}

func Encode(w io.Writer, im image.Image) error {
	b := im.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, im, b.Min, draw.Src)

	id := newIcondir()
	ide := newIcondirentry()

	pngbuffer := new(bytes.Buffer)
	pngwriter := bufio.NewWriter(pngbuffer)
	err := png.Encode(pngwriter, m)
	if err != nil {
		return err
	}
	err = pngwriter.Flush()
	if err != nil {
		return err
	}
	ide.Size = uint32(len(pngbuffer.Bytes()))

	bounds := m.Bounds()
	ide.Width = uint8(bounds.Dx())
	ide.Height = uint8(bounds.Dy())
	bb := new(bytes.Buffer)

	var e error
	e = binary.Write(bb, binary.LittleEndian, id)
	if e != nil {
		return e
	}
	e = binary.Write(bb, binary.LittleEndian, ide)
	if e != nil {
		return e
	}

	_, e = w.Write(bb.Bytes())
	if e != nil {
		return e
	}
	_, e = w.Write(pngbuffer.Bytes())
	if e != nil {
		return e
	}

	return e
}
