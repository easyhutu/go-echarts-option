package tests

import (
	"encoding/json"
	"github.com/easyhutu/go-echarts-option/model/opts/grid"
	"github.com/easyhutu/go-echarts-option/model/types/colors"
	"testing"
)

func TestName(t *testing.T) {
	g := grid.Grid{
		ContainLabel: true,
		BorderColor:  colors.NewImg("/src/test.img"),
	}
	h, _ := json.Marshal(g)
	println(string(h))
}
