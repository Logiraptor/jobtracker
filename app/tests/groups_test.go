package tests

import (
	"github.com/golang/mock/gomock"

	"testing"
)

type TestInterface interface {
	Before(string)
	After(string)
	Body(string)
}

func TestDescribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ti := NewMockTestInterface(ctrl)
	gomock.InOrder(
		ti.EXPECT().Before("A"),
		ti.EXPECT().Body("C"),
		ti.EXPECT().After("B"),

		ti.EXPECT().Before("A"),
		ti.EXPECT().Before("D"),
		ti.EXPECT().Body("F"),
		ti.EXPECT().After("E"),
		ti.EXPECT().After("B"),

		ti.EXPECT().Before("A"),
		ti.EXPECT().Before("D"),
		ti.EXPECT().Body("G"),
		ti.EXPECT().After("E"),
		ti.EXPECT().After("B"),

		ti.EXPECT().Before("A"),
		ti.EXPECT().Before("D"),
		ti.EXPECT().Body("H"),
		ti.EXPECT().After("E"),
		ti.EXPECT().After("B"),
	)

	Describe(t, "Outer Describe", func(c *Context) {
		c.Before(func() {
			ti.Before("A")
		})
		c.After(func() {
			ti.After("B")
		})

		c.It("It 1", func() {
			ti.Body("C")
		})

		c.Describe("Inner Describe", func(c *Context) {
			c.Before(func() {
				ti.Before("D")
			})
			c.After(func() {
				ti.After("E")
			})
			c.It("It 2", func() {
				ti.Body("F")
			})
			c.It("It 3", func() {
				ti.Body("G")
			})
			c.It("It 4", func() {
				ti.Body("H")
			})
		})
	})
}
