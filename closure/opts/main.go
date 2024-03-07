package main

import "fmt"

type Persion struct {
	Name   string
	Age    int
	Talk   func(content string)
	Listen func()
}

func NewPersion(opts ...Option) *Persion {
	p := &Persion{}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

type Option func(p *Persion)

func OptName(name string) Option {
	return func(p *Persion) {
		p.Name = name
	}
}

func OptAge(age int) Option {
	return func(p *Persion) {
		p.Age = age
	}
}

func OptTalk(talk func(string)) Option {
	return func(p *Persion) {
		p.Talk = talk
	}
}

func OptListen(listen func()) Option {
	return func(p *Persion) {
		p.Listen = listen
	}
}

func main() {
	// p1 := NewPersion(OptName("Zhao"))
	// fmt.Println(p1.Name)

	// p2 := NewPersion(OptName("Li"), OptAge(26))
	// fmt.Println(p2)

	ch := make(chan string)
	talk := func(content string) {
		ch <- content
	}
	listen := func() {
		for content := range ch {
			if content == "bye" {
				break
			}
			fmt.Println(content)
		}
	}
	p3 := NewPersion(OptTalk(talk), OptAge(18))
	p4 := NewPersion(OptListen(listen))
	go func() {
		p3.Talk("hello")
		p3.Talk("I'm " + p3.Name)
		p3.Talk(fmt.Sprintf("I'm %d", p3.Age))
		p3.Talk("bye")
	}()

	p4.Listen()
}
