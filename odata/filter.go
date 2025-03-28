package odata

import "fmt"

func NewFilter() *Filter {
	return &Filter{}
}

type Filter struct {
	Query string
}

func (f *Filter) Set(q string) {
	f.Query = q
}

func (f *Filter) MarshalSchema() string {
	return f.Query
}

func (f Filter) IsZero() bool {
	return f.Query == ""
}

func (t Filter) Format(f fmt.State, c rune) {
	f.Write([]byte(t.MarshalSchema()))
}
