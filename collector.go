package govaluate

import (
	"errors"
	"fmt"
)

type Result struct {
	Expression string
	Evaluation bool
}

type Collector struct {
	Params  *ParamBag
	Results []Result
}

func NewCollector() *Collector {
	return &Collector{Params: NewCollectingParams()}
}

func (c *Collector) Decorate(stage *evaluationStage) {
	// TODO(Dean): Comparison operation map instead
	if stageSymbolMap[stage.symbol] != nil {
		c.Results = append(c.Results, Result{
			Expression: fmt.Sprintf("%s %s %s", c.unpack(stage.leftStage), stage.symbol, c.unpack(stage.rightStage)),
			Evaluation: false, // todo??
		})
	}
}

func (c *Collector) Get(name string) (interface{}, error) {
	return c.Params.Get(name)
}
func (c *Collector) unpack(stage *evaluationStage) string {
	lastParam := c.Params.LastParam()
	l, _ := stage.operator(nil, nil, c.Params)
	if lastParam != c.Params.LastParam() {
		return c.Params.LastParam()
	}

	return fmt.Sprint(l)
}

type ParamBag struct {
	params    MapParameters
	collected []string
}

func NewCollectingParams() *ParamBag {
	return &ParamBag{params: make(map[string]interface{})}
}

func (c *ParamBag) Set(name string, value interface{}) {
	c.params[name] = value
}

func (c *ParamBag) Get(name string) (interface{}, error) {
	v, ok := c.params[name]
	if !ok {
		return nil, errors.New("value doesnt exist")
	}
	c.collected = append(c.collected, name)
	return v, nil
}

func (c *ParamBag) LastParam() string {
	if len(c.collected) == 0 {
		return ""
	}
	return c.collected[len(c.collected)-1]
}
