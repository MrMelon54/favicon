package favicon

import (
	"bytes"
	"code.mrmelon54.xyz/sean/png2ico"
	_ "embed"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"github.com/tdewolff/canvas/renderers/svg"
	"image/color"
	"strings"
)

//go:embed Lato-Bold.ttf
var latoBold []byte

type Svg struct {
	c      *canvas.Canvas
	color  int
	letter rune
}

func NewSvg(addr string, faviconColor *Color) *Svg {
	return &Svg{canvas.New(96, 96), faviconColor.PickColor(addr), []rune(strings.ToUpper(addr))[0]}
}

func (favicon *Svg) generate() {
	favicon.c.Reset()
	ctx := canvas.NewContext(favicon.c)
	round := canvas.RoundedRectangle(96, 96, 8)
	ctx.SetFillColor(color.RGBA{R: favicon.getR(), G: favicon.getG(), B: favicon.getB(), A: 0xff})
	ctx.DrawPath(0, 0, round)

	fontFamily := canvas.NewFontFamily("Lato")
	if err := fontFamily.LoadFont(latoBold, 0, canvas.FontBold); err != nil {
		panic(err)
	}

	face := fontFamily.Face(100, canvas.White, canvas.FontBold, canvas.FontNormal)
	fontPath, _, err := face.ToPath(string(favicon.letter))
	if err != nil {
		panic(err)
	}

	ctx.SetFillColor(color.White)
	ctx.Translate(16, -10)
	ctx.Rotate(10)
	ctx.Scale(3.6, 3.6)
	ctx.DrawPath(0, 0, fontPath)
}

func (favicon *Svg) getR() uint8 { return uint8(favicon.color >> 16) }
func (favicon *Svg) getG() uint8 { return uint8(favicon.color >> 8) }
func (favicon *Svg) getB() uint8 { return uint8(favicon.color) }

func (favicon *Svg) ProduceSvg() ([]byte, error) {
	favicon.generate()
	b := new(bytes.Buffer)
	err := renderers.SVG(&svg.Options{EmbedFonts: false, SubsetFonts: false, ImageEncoding: canvas.Lossless})(b, favicon.c)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (favicon *Svg) ProducePng() ([]byte, error) {
	favicon.generate()
	b := new(bytes.Buffer)
	err := renderers.PNG()(b, favicon.c)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (favicon *Svg) ProduceIco() ([]byte, error) {
	b, err := favicon.ProducePng()
	if err != nil {
		return nil, err
	}
	return png2ico.ConvertPngToIco(b, int(favicon.c.W), int(favicon.c.H))
}
