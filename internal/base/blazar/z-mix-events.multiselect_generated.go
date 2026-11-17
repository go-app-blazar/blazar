package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarMultiselect) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarMultiselect {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarMultiselect) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarMultiselect {
	c.MixOnClick(function)
	return c
}

func (c *blazarMultiselect) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarMultiselect {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarMultiselect) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarMultiselect {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarMultiselect) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarMultiselect {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarMultiselect) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarMultiselect {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarMultiselect) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarMultiselect {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarMultiselect) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarMultiselect {
	c.MixOnKeyUp(function)
	return c
}
