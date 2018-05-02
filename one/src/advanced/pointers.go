package advanced

import (
	"strconv"
)

type (
	Fn func(string) int
	Fi func(int) string
)

var (
	ToInt Fn = func(str string) int {
		i, _ := strconv.Atoi(str)
		return i
	}

	ToStr Fi = func(i int) string {
		s := strconv.Itoa(i)
		return s
	}
)

func Ex(i int) bool {
	return i == ToInt(ToStr(i))
}

type A struct {
	Name string
}

func (a *A) Rename(b *A) (*A, error) {
	a.Name = b.Name
	c := *a
	return &c, nil
}
