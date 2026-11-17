package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarIcon) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarIcon {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarIcon) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarIcon {
	c.MixOnClick(function)
	return c
}

func (c *blazarIcon) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarIcon {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarIcon) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarIcon {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarIcon) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarIcon {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarIcon) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarIcon {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarIcon) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarIcon {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarIcon) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarIcon {
	c.MixOnKeyUp(function)
	return c
}
