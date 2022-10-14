package test

type Foo struct {
	bar1 string
	Bar2 string
}

func NewFoo(bar1, bar2 string) *Foo {
	return &Foo{
		bar1: bar1,
		Bar2: bar2,
	}
}
