package factorymethod

type Operator interface {
	SetA(int)
	SetB(int)
	GetResult() int
}

type OperatorFactory interface {
	CreateOperator() Operator
}

type OperatorBase struct {
	a int
	b int
}

func (op *OperatorBase) SetA(a int) {
	op.a = a
}

func (op *OperatorBase) SetB(b int) {
	op.b = b
}

type AddOperatorFactory struct{}

type AddOperator struct {
	*OperatorBase
}

func (op *AddOperator) GetResult() int {
	return op.a + op.b
}

func (AddOperatorFactory) CreateOperator() Operator {
	return &AddOperator{OperatorBase: &OperatorBase{}}
}

type MinusOperatorFactory struct{}

func (MinusOperatorFactory) CreateOperator() Operator {
	return &MinusOperator{OperatorBase: &OperatorBase{}}
}

type MinusOperator struct {
	*OperatorBase
}

func (op *MinusOperator) GetResult() int {
	return op.a - op.b
}
