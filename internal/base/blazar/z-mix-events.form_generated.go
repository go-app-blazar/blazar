package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarForm) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarForm {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarForm) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarForm {
	c.MixOnClick(function)
	return c
}

func (c *blazarForm) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarForm {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarForm) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarForm {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarForm) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarForm {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarForm) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarForm {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarForm) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarForm {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarForm) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarForm {
	c.MixOnKeyUp(function)
	return c
}
