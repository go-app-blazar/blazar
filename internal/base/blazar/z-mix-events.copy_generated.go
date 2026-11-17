package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarCopy) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarCopy {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarCopy) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarCopy {
	c.MixOnClick(function)
	return c
}

func (c *blazarCopy) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarCopy {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarCopy) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarCopy {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarCopy) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarCopy {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarCopy) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarCopy {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarCopy) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarCopy {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarCopy) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarCopy {
	c.MixOnKeyUp(function)
	return c
}
