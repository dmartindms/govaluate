package govaluate

import "errors"

type CollectingParamBag struct {
	params    MapParameters
	collected []string
}

func NewCollectingParams() *CollectingParamBag {
	return &CollectingParamBag{params: make(map[string]interface{})}
}

func (c *CollectingParamBag) Set(name string, value interface{}) {
	c.params[name] = value
}

func (c *CollectingParamBag) Get(name string) (interface{}, error) {
	v, ok := c.params[name]
	if !ok {
		return nil, errors.New("value doesnt exist")
	}
	c.collected = append(c.collected, name)
	return v, nil
}

func (c *CollectingParamBag) LastParam() string {
	if len(c.collected) == 0 {
		return ""
	}
	return c.collected[len(c.collected)-1]
}
