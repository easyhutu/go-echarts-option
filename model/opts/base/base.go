package base

import (
	"github.com/easyhutu/go-echarts-option/model/types"
)

type Sizes struct {
	Left   types.Size `json:"left,omitempty"`
	Top    types.Size `json:"top,omitempty"`
	Right  types.Size `json:"right,omitempty"`
	Bottom types.Size `json:"bottom,omitempty"`
	Width  types.Size `json:"width,omitempty"`
	Height types.Size `json:"height,omitempty"`
}

type TextStyle struct {
	Color           types.Color `json:"color,omitempty"`
	FontStyle       string      `json:"fontStyle,omitempty"`
	FontWidget      string      `json:"fontWidget,omitempty"`
	FontFamily      string      `json:"fontFamily,omitempty"`
	FontSize        string      `json:"fontSize,omitempty"`
	Align           string      `json:"align,omitempty"`
	BackgroundColor types.Color `json:"backgroundColor,omitempty"`
	BorderColor     types.Color `json:"borderColor,omitempty"`
}
