package inject

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

type InterfaceA interface {
	MethodA()
	MethodB()
}

type InterfaceB interface {
	MethodC()
	MethodD()
}

type interfaceAImpl struct{}

func (i *interfaceAImpl) MethodA() {}
func (i *interfaceAImpl) MethodB() {}

type interfaceBImpl struct{}

func (i *interfaceBImpl) MethodC() {}
func (i *interfaceBImpl) MethodD() {}

type ImplementationA struct{ interfaceAImpl }

type ImplementationB struct {
	interfaceAImpl
	id int
}

type DownstreamType struct {
	interfaceBImpl
	InterfaceA
}

type UnregisteredInterface1 interface {
	MethodX()
}

type UnregisteredInterface2 interface {
	MethodX()
}

func TestInject(t *testing.T) {
	c := NewContainer()

	c.Register((*InterfaceA)(nil), &ImplementationA{})
	var i InterfaceA
	assert.True(t, c.Resolve(&i))
	assert.IsType(t, &ImplementationA{}, i)

	type Struct struct {
		InterfaceA
	}

	var s Struct
	assert.Nil(t, c.FillStruct(&s))

	assert.IsType(t, &ImplementationA{}, s.InterfaceA)

	c.RegisterContructor((*InterfaceA)(nil), func() func() interface{} {
		i := 0
		return func() interface{} {
			i += 1
			return &ImplementationB{id: i}
		}
	}())

	for n := 1; n < 10; n++ {
		c.Resolve(&i)
		assert.IsType(t, &ImplementationB{}, i)
		assert.Equal(t, n, i.(*ImplementationB).id)
	}

	var unregistered UnregisteredInterface1
	assert.False(t, c.Resolve(&unregistered))
	assert.Nil(t, unregistered)

	type ImpossibleStruct struct {
		InterfaceA
		UnregisteredInterface1
		UnregisteredInterface2
	}

	var is ImpossibleStruct
	err := c.FillStruct(&is)
	assert.Contains(t, err.Error(), "UnregisteredInterface1")
	assert.Contains(t, err.Error(), "UnregisteredInterface2")
	assert.Contains(t, err.Error(), "ImpossibleStruct")
	assert.Len(t, err, 2)

	c.Register((*InterfaceB)(nil), &DownstreamType{})
	type App struct {
		InterfaceB
	}
	var app App
	err = c.FillStruct(&app)
	assert.Nil(t, err)
	assert.IsType(t, &DownstreamType{}, app.InterfaceB)
	assert.IsType(t, &ImplementationB{}, app.InterfaceB.(*DownstreamType).InterfaceA)
}
