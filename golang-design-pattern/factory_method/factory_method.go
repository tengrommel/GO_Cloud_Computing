package factory_method

// Operator是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase是Operator接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

// SetA 设置　A
func (o *OperatorBase)SetA(a int)  {
	o.a = a
}

func (o *OperatorBase)SetB(b int)  {
	o.b = b
}

// PlusOperatorFactory是PlusOperator的工厂类
type PlusOperatorFactory struct {}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// PlusOperator Operator的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o PlusOperator)Result() int {
	return o.a + o.b
}

// MinusOperatorFactor是MinusOperator的工厂类
type MinusOperatorFactory struct {}

func (MinusOperatorFactory)Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperator Operator的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o MinusOperator)Result() int {
	return o.a - o.b
}