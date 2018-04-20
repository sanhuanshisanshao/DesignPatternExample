package main

//定义实物的接口
type Food interface {
	GetDescription() string
	SetDescription(s string)
	GetCost() float32
	SetCost(c float32)
}

type inter interface {
	getDesc() string
	getCosts() float32
}

//装饰器
type decorator func(food Food) Food

type mocha struct {
	desc string
	cost float32
}

func (m *mocha) getDesc() string {
	return m.desc
}

func (m *mocha) getCosts() float32 {
	return m.cost
}

type roast struct {
	desc string
	cost float32
}

func (r *roast) getDesc() string {
	return r.desc
}

func (r *roast) getCosts() float32 {
	return r.cost
}

func mochaDecorator(m mocha) decorator {
	return func(f Food) Food {
		f.SetCost(f.GetCost() + m.getCosts())
		f.SetDescription(f.GetDescription() + m.getDesc())
		return f
	}
}

func roastDecorator(r roast) decorator {
	return func(food Food) Food {
		food.SetDescription(food.GetDescription() + r.getDesc())
		food.SetCost(food.GetCost() + r.getCosts())
		return food
	}
}

func decorators(f inter, ds ...decorator) inter {
	for k, v := range ds {

	}
	return f
}

func main() {

}
