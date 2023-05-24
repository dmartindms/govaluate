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

	params := NewCollectingParams()
	params.Set("age", 21)
	res, err := exp.Eval(params)

	if err != nil {
		t.Error(err)
	}
	log.Println(exp, "==", res)

	if params.LastParam() != "age" {
		t.Error("failed to log last parameter, got: " + params.LastParam())
	}

	var results = make(map[string]bool)
	if results["age >= 21"] != true {
		t.Error("result 'age >= 21' was not in the results map")
	}
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
