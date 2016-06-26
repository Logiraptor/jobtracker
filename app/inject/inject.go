package inject

import (
	"fmt"
	"reflect"
	"strings"
)

type Container struct {
	available map[reflect.Type]resolver
}

func NewContainer() *Container {
	return &Container{
		available: make(map[reflect.Type]resolver),
	}
}

func (c *Container) Register(ptrType, impl interface{}) {
	val := reflect.ValueOf(impl)
	if val.Elem().Kind() == reflect.Struct {
		c.FillStruct(impl)
	}
	c.available[reflect.TypeOf(ptrType).Elem()] = staticResolver{
		impl: reflect.ValueOf(impl),
	}
}

func (c *Container) RegisterContructor(ptrType interface{}, cons func() interface{}) {
	c.available[reflect.TypeOf(ptrType).Elem()] = constructorResolver{
		constructor: cons,
	}
}

func (c *Container) Resolve(target interface{}) bool {
	if resolver, ok := c.available[reflect.TypeOf(target).Elem()]; ok {
		resolver.Resolve(reflect.ValueOf(target).Elem())
		return true
	}
	return false
}

func (c *Container) FillStruct(s interface{}) error {
	val := reflect.ValueOf(s).Elem()
	var errs MultiError
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if val.Type().Field(i).PkgPath != "" {
			continue
		}
		if !c.Resolve(field.Addr().Interface()) {
			errs = append(errs, fmt.Errorf("Could not resolve type %q while filling %q", val.Type().Name(), field.Type().Name()))
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

type MultiError []error

func (m MultiError) Error() string {
	var errs = make([]string, len(m))
	for i, e := range m {
		errs[i] = e.Error()
	}
	return strings.Join(errs, "\n")
}

type resolver interface {
	Resolve(target reflect.Value)
}

type staticResolver struct {
	impl reflect.Value
}

func (s staticResolver) Resolve(target reflect.Value) {
	target.Set(s.impl)
}

type constructorResolver struct {
	constructor func() interface{}
}

func (c constructorResolver) Resolve(target reflect.Value) {
	target.Set(reflect.ValueOf(c.constructor()))
}
