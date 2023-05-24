package govaluate

import (
	"fmt"
	"log"
	"testing"
)

func TestSimpleResultCollection(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	exp, err := NewEvaluableExpression(`age >= 21`)
	if err != nil {
		t.Error("couldn't parse expression: ", err)
	}

	collector := NewCollector()
	collector.Params.Set("age", 21)
	collector.Decorate(exp.evaluationStages)
	if collector.Results[0].Expression != "age >= 21" {
		t.Error("failed to collect results")
	}
	fmt.Println(collector.Results)

	res, err := exp.Eval(collector)
	if err != nil {
		t.Error(err)
	}
	log.Println(exp, "==", res)
}

func TestDecoratorCollector(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	exp, err := NewEvaluableExpression(`age >= 21 && age <= 24 && state IN ("GA", "FL")`)
	if err != nil {
		t.Error("couldn't parse expression: ", err)
	}

	fmt.Println(exp.evaluationStages.symbol)

	params := make(MapParameters)
	params["age"] = 21
	params["state"] = "GA"
	res, err := exp.Eval(params)

	if err != nil {
		t.Error(err)
	}
	log.Println(exp, "==", res)
}
