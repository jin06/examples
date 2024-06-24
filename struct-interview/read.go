package main

type Animal interface {
	Run() string
}

type Dog struct {
}

func (*Dog) Run() string {
	return "dog run"
}

// var _ =  (Animal)((*Dog)(nil))
var _ = (Animal)((*Dog)(nil))
