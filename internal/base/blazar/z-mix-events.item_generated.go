package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarItem) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarItem {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarItem) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarItem {
	c.MixOnClick(function)
	return c
}

func (c *blazarItem) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarItem {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarItem) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarItem {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarItem) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarItem {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarItem) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarItem {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarItem) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarItem {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarItem) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarItem {
	c.MixOnKeyUp(function)
	return c
}
