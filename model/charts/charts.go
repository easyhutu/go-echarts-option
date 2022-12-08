package charts

import (
	"github.com/easyhutu/go-echarts-option/model/charts/line"
	"github.com/easyhutu/go-echarts-option/model/opts/titles"
)

type (
	Line = line.Line
)

type Chart struct {
	Title titles.Title `json:"title,omitempty"`
}
