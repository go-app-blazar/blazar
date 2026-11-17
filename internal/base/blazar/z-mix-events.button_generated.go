package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarButton) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarButton {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarButton) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarButton {
	c.MixOnClick(function)
	return c
}

func (c *blazarButton) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarButton {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarButton) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarButton {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarButton) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarButton {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarButton) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarButton {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarButton) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarButton {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarButton) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarButton {
	c.MixOnKeyUp(function)
	return c
}
