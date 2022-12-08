package series

import (
	"github.com/easyhutu/go-echarts-option/model/opts"
	"github.com/easyhutu/go-echarts-option/model/types"
)

type Series struct {
	Type            opts.Type   `json:"type,omitempty"`
	Name            string      `json:"name,omitempty"`
	ColorBy         string      `json:"colorBy,omitempty"` // series, data
	XAxisIndex      int         `json:"xAxisIndex,omitempty"`
	YAxisIndex      int         `json:"yAxisIndex,omitempty"`
	PolarIndex      int         `json:"polarIndex,omitempty"`
	Symbol          string      `json:"symbol,omitempty"` //circle, rect, roundRect, triangle, diamond, pin, arrow, none
	SymbolSize      int         `json:"symbolSize,omitempty"`
	SymbolRotate    int         `json:"symbolRotate,omitempty"`
	LegendHoverLink bool        `json:"legendHoverLink,omitempty"`
	Stack           string      `json:"stack,omitempty"`
	Cursor          string      `json:"cursor,omitempty"` // pointer,help
	ConnectNulls    bool        `json:"connectNulls,omitempty"`
	Clip            bool        `json:"clip,omitempty"`
	Step            bool        `json:"step,omitempty"`
	Label           Label       `json:"label,omitempty"`
	Data            interface{} `json:"data,omitempty"`
}

type Label struct {
	Position  types.Position `json:"position,omitempty"`
	Formatter string         `json:"formatter,omitempty"` // formatter: '{b}: {@score}'
	Color     types.Color    `json:"color,omitempty"`
	FontStyle string         `json:"fontStyle,omitempty"` // normal, italic, oblique
}
