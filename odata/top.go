package odata

import (
	"fmt"
	"strconv"
)

func NewTop() *Top {
	return &Top{}
}

type Top struct {
	i int
}

func (t *Top) Set(i int) {
	t.i = i
}

func (t *Top) MarshalSchema() string {
	i := int64(t.i)
	if i == 0 {
		return ""
	}

	return strconv.FormatInt(i, 10)
}

func (t Top) IsZero() bool {
	return t.i == 0
}

func (t Top) Format(f fmt.State, c rune) {
	f.Write([]byte(t.MarshalSchema()))
}
