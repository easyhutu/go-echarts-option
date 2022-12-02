package colors

import "fmt"

type Color interface {
	Show()
}

type RGBA string

func (r RGBA) Show() {
	println(r)
}

func NewRGBA(a, b, c int, d float32) Color {
	if d == 0 {
		return RGBA(fmt.Sprintf("rgb(%v, %v, %v)", a, b, c))
	}
	return RGBA(fmt.Sprintf("rgba(%d, %d, %d, %f)", a, b, c, d))
}

type Img struct {
	Image string `json:"image,omitempty"`
}

func (i *Img) Show() {
	println(i)
}

func NewImg(src string) Color {
	return &Img{src}
}

type Hex string

func (h Hex) Show() {
	println(h)
}

func NewHex(hex string) Color {
	h := Hex(hex)
	return h
}

var (
	C1 = NewHex("#5470c6")
	C2 = NewHex("#91cc75")
	C3 = NewHex("#fac858")
	C4 = NewHex("#ee6666")
	C5 = NewHex("#73c0de")
	C6 = NewHex("#3ba272")
	C7 = NewHex("#fc8452")
	C8 = NewHex("#9a60b4")
	C9 = NewHex("#ea7ccc")
)
