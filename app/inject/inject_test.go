package inject

import (
	"jobtracker/app/tests"

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
	tests.Describe(t, "Container", func(c *tests.Context) {
		var container *Container
		c.Before(func() {
			container = NewContainer()
		})

		c.Describe("Given I do not register a dependency", func(c *tests.Context) {
			c.It("Fails to resolve", func() {
				var unregistered UnregisteredInterface1
				assert.False(t, container.Resolve(&unregistered))
				assert.Nil(t, unregistered)
			})

			c.Describe("When I fill a struct", func(c *tests.Context) {
				c.It("Returns a useful error message", func() {
					type ImpossibleStruct struct {
						InterfaceA
						UnregisteredInterface1
						UnregisteredInterface2
					}

					var is ImpossibleStruct
					err := container.FillStruct(&is)
					assert.Contains(t, err.Error(), "UnregisteredInterface1")
					assert.Contains(t, err.Error(), "UnregisteredInterface2")
					assert.Contains(t, err.Error(), "ImpossibleStruct")
					assert.Len(t, err, 3)
				})
			})
		})

		c.Describe("Given I register a dependency", func(c *tests.Context) {
			c.Before(func() {
				container.Register((*InterfaceA)(nil), &ImplementationA{})
			})

			c.It("Can be resolved", func() {
				var i InterfaceA
				assert.True(t, container.Resolve(&i))
				assert.IsType(t, &ImplementationA{}, i)
			})

			c.It("Can be resolved into a struct", func() {
				type Struct struct {
					InterfaceA
				}

				var s Struct
				assert.Nil(t, container.FillStruct(&s))
				assert.IsType(t, &ImplementationA{}, s.InterfaceA)
			})

			c.It("Autofills structs I register later", func() {
				container.Register((*InterfaceB)(nil), &DownstreamType{})
				type App struct {
					InterfaceB
				}
				var app App
				err := container.FillStruct(&app)
				assert.Nil(t, err)
				assert.IsType(t, &DownstreamType{}, app.InterfaceB)
				assert.IsType(t, &ImplementationA{}, app.InterfaceB.(*DownstreamType).InterfaceA)
			})
		})

		c.Describe("Given I register a constructor", func(c *tests.Context) {
			c.Before(func() {
				container.RegisterContructor((*InterfaceA)(nil), func() func() interface{} {
					i := 0
					return func() interface{} {
						i += 1
						return &ImplementationB{id: i}
					}
				}())
			})

			c.It("Resolves into a new instance each time", func() {
				var i InterfaceA
				for n := 1; n < 10; n++ {
					container.Resolve(&i)
					assert.IsType(t, &ImplementationB{}, i)
					assert.Equal(t, n, i.(*ImplementationB).id)
				}
			})
		})
	})
}
