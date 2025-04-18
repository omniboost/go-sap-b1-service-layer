package odata

import (
	"fmt"
	"strconv"
)

func NewCount() *Count {
	return &Count{}
}

type Count struct {
	i int
}

func (t *Count) Set(i int) {
	t.i = i
}

func (t *Count) MarshalSchema() string {
	i := int64(t.i)
	if i == 0 {
		return ""
	}

	return strconv.FormatInt(i, 10)
}

func (t Count) IsZero() bool {
	return t.i == 0
}

func (t Count) Format(f fmt.State, c rune) {
	f.Write([]byte(t.MarshalSchema()))
}
