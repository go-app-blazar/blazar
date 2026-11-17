package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarInput[T]) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarInput[T] {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarInput[T]) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarInput[T] {
	c.MixOnClick(function)
	return c
}

func (c *blazarInput[T]) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarInput[T] {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarInput[T]) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarInput[T] {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarInput[T]) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarInput[T] {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarInput[T]) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarInput[T] {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarInput[T]) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarInput[T] {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarInput[T]) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarInput[T] {
	c.MixOnKeyUp(function)
	return c
}
