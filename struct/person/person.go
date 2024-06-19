package person

type GrandFather struct{}

func (g *GrandFather) Echo(s string) string {
	return s
}

func (g *GrandFather) Hello() string {
	return "grand fahter say hello"
}

type Father struct {
	SecondName string
}

func (f *Father) Hello() string {
	return "father say hello"
}

// 子类
type Son struct {
	Father
	GrandFather
}

func (s *Son) Hello2() string {
	return "son say hello"
}
