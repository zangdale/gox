// https://pkg.go.dev/github.com/nfnt/resize
package resize

import (
	"image"

	r "github.com/nfnt/resize"
)

// Resize 使用插值函数 interp 将图像缩放到新的宽度和高度。将返回具有给定尺寸的新图像。
// 如果参数宽度或高度之一设置为 0，则将计算其大小，以便纵横比是原始图像的纵横比。
// 调整大小算法使用通道进行并行计算。如果输入图像的宽度或高度为 0，则原样返回。
func Resize(width, height uint, img image.Image, interp r.InterpolationFunction) image.Image {
	return r.Resize(width, height, img, interp)
}

// Thumbnail 将提供的图像缩小到最大宽度和高度，保留原始纵横比并使用插值函数 interp。
// 如果原始尺寸已经小于提供的约束，它将返回原始图像，而不对其进行处理。
func Thumbnail(maxWidth, maxHeight uint, img image.Image, interp r.InterpolationFunction) image.Image {
	return r.Thumbnail(maxWidth, maxHeight, img, interp)
}
