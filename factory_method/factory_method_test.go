package factorymethod

import (
	"log"
	"testing"
)

func compute(factory OperatorFactory, a, b int) int {
	operator := factory.CreateOperator()
	operator.SetA(a)
	operator.SetB(b)
	return operator.GetResult()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)

	log.Printf("plus test")
	factory = &AddOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		t.Error("AddOperatorFactory failed")
	}

	factory = &MinusOperatorFactory{}
	if compute(factory, 5, 3) != 2 {
		t.Error("MinusOperatorFactory failed")
	}
}
