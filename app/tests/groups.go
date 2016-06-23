package tests

import (
	"strings"
	"testing"
)

type tab int

func (t tab) String() string {
	return strings.Repeat("\t", int(t))
}

type describe struct {
	name string
	body func(*Context)
}

type it struct {
	name string
	body func()
}

type Context struct {
	depth     tab
	describes []describe
	befores   []func()
	afters    []func()
	its       []it
}

func Describe(t *testing.T, name string, body func(c *Context)) {
	c := &Context{
		depth:     0,
		describes: []describe{{name: name, body: body}},
	}
	c.run(t)
}

func (c *Context) run(t *testing.T) {
	for _, i := range c.its {
		t.Logf("%sIt %s", c.depth, i.name)
		c.setup()
		i.body()
		c.teardown()
	}

	for _, d := range c.describes {
		t.Logf("%sDescribe %s", c.depth, d.name)
		subContext := &Context{
			depth:   c.depth + 1,
			befores: c.befores[:len(c.befores):len(c.befores)],
			afters:  c.afters[:len(c.afters):len(c.afters)],
		}
		d.body(subContext)
		subContext.run(t)
	}
}

func (c *Context) setup() {
	for _, beforeFunc := range c.befores {
		beforeFunc()
	}
}

func (c *Context) teardown() {
	for i := len(c.afters) - 1; i >= 0; i-- {
		c.afters[i]()
	}
}

func (c *Context) Describe(name string, body func(c *Context)) {
	c.describes = append(c.describes, describe{name: name, body: body})
}

func (c *Context) It(name string, body func()) {
	c.its = append(c.its, it{name: name, body: body})
}

func (c *Context) Before(body func()) {
	c.befores = append(c.befores, body)
}

func (c *Context) After(body func()) {
	c.afters = append(c.afters, body)
}
