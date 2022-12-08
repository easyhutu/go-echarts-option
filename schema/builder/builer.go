package builder

import "github.com/easyhutu/go-echarts-option/schema"

type Base interface {
	SetTitle(ctx schema.Context) Base
	SetTooltip() Base
	GetChart() map[string]interface{}
}

type base struct {
	Chart map[string]interface{}
}

func (b base) GetChart() map[string]interface{} {
	return b.Chart
}

func (b base) SetTitle(ctx schema.Context) Base {
	b.Chart["title"] = ctx.Title()
	return b
}

func (b base) SetTooltip() Base {
	panic("")
}

func NewBaseBuilder() Base {
	return &base{}
}
