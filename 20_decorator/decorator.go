package decorator

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

//装饰器， 而非简单工厂， 注意返回值
func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

//此处 改变对象功能， 使之返回 `num` 倍的`Calc()`
func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}


type AddDecorator struct {
	Component
	num int
}

//此处 改变对象功能
func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
