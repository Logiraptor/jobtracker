package doubles

import (
	"fmt"

	"net/http"
)

type FakeView struct {
	Render_ func(http.ResponseWriter, string, interface{}) error
}

func NewFakeView() *FakeView {
	return &FakeView{
		Render_: func(rw http.ResponseWriter, name string, data interface{}) error {
			fmt.Fprintf(rw, "name: %s, data: %#v", name, data)
			return nil
		},
	}
}

func (f *FakeView) Render(rw http.ResponseWriter, name string, data interface{}) error {
	return f.Render_(rw, name, data)
}
