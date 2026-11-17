package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarInputWrapper) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarInputWrapper {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarInputWrapper) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarInputWrapper {
	c.MixOnClick(function)
	return c
}

func (c *blazarInputWrapper) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarInputWrapper {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarInputWrapper) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarInputWrapper {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarInputWrapper) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarInputWrapper {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarInputWrapper) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarInputWrapper {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarInputWrapper) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarInputWrapper {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarInputWrapper) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarInputWrapper {
	c.MixOnKeyUp(function)
	return c
}
