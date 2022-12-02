package legend

import (
	"github.com/easyhutu/go-echarts-option/model/opts/base"
)

const (
	TypePlain  Type = "plain"
	TypeScroll Type = "scroll"
)

// Legend 图例组件。https://echarts.apache.org/zh/option.html#legend
type Legend struct {
	Type Type `json:"type,omitempty"`
	base.Sizes
}

type Type string
