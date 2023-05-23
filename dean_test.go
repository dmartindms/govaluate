package govaluate

import (
	"fmt"
	"log"
	"testing"
)

var counter = 0

func decorate(stage *evaluationStage) {
}

func TestDecoratorCollector(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	exp, err := NewEvaluableExpression(`age >= 21 && age <= 24 && state IN ("GA", "FL")`)
	if err != nil {
		t.Error("couldn't parse expression: ", err)
	}

	fmt.Println(exp.evaluationStages.symbol)
	decorate(exp.evaluationStages)

	params := make(MapParameters)
	params["age"] = 21
	params["state"] = "GA"
	res, err := exp.Eval(params)

	if err != nil {
		t.Error(err)
	}
	log.Println(exp, "==", res)
}
