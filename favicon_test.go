package favicon

import (
	"bytes"
	_ "embed"
	"testing"
)

var (
	//go:embed test_images/example.svg
	exampleFaviconSvg string
	//go:embed test_images/example.png
	exampleFaviconPng []byte
	//go:embed test_images/example.ico
	exampleFaviconIco []byte
)

func TestFaviconSvg(t *testing.T) {
	c := NewColor()
	f := NewSvg("example.com", c)
	b, err := f.ProduceSvg()
	if err != nil {
		t.Fatalf("Failed to generate the SVG: %s", err)
	}
	if string(b) != exampleFaviconSvg {
		t.Fatalf("Generated SVG doesn't match expected result, got: %s", string(b))
	}
}

func TestFaviconPng(t *testing.T) {
	c := NewColor()
	f := NewSvg("example.com", c)
	b, err := f.ProducePng()
	if err != nil {
		t.Fatalf("Failed to rasterize the PNG: %s", err)
	}
	if bytes.Compare(b, exampleFaviconPng) != 0 {
		t.Fatalf("Generated PNG doesn't match expected result")
	}
}

func TestFaviconIco(t *testing.T) {
	c := NewColor()
	f := NewSvg("example.com", c)
	b, err := f.ProduceIco()
	if err != nil {
		t.Fatalf("Failed to rasterize the ICO: %s", err)
	}
	if bytes.Compare(b, exampleFaviconIco) != 0 {
		t.Fatalf("Generated ICO doesn't match expected result")
	}
}
