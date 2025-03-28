package odata

import (
	"fmt"
	"strconv"
)

func NewSkip() *Skip {
	return &Skip{}
}

type Skip struct {
	i int
}

func (t *Skip) Set(i int) {
	t.i = i
}

func (t *Skip) MarshalSchema() string {
	i := int64(t.i)
	if i == 0 {
		return ""
	}

	return strconv.FormatInt(i, 10)
}

func (t Skip) IsZero() bool {
	return t.i == 0
}

func (t Skip) Format(f fmt.State, c rune) {
	f.Write([]byte(t.MarshalSchema()))
}
