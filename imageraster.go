package snowfavicon

import (
	"bytes"
	"code.mrmelon54.xyz/sean/png2ico"
	"code.mrmelon54.xyz/sean/svg2png"
)

func RasterSvgToPng(d *bytes.Buffer) ([]byte, error) {
	png, err := svg2png.RasterSvgToPng(bytes.NewReader(d.Bytes()))
	if err != nil {
		return nil, err
	}
	return png, nil
}

func RasterSvgToIco(d *bytes.Buffer) ([]byte, error) {
	b, err := RasterSvgToPng(d)
	if err != nil {
		return nil, err
	}
	return png2ico.ConvertPngToIco(b, 96, 96)
}
