package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func (c *blazarSelect) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarSelect {
	c.MixOn(name, eventHandler, options...)
	return c
}

func (c *blazarSelect) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarSelect {
	c.MixOnClick(function)
	return c
}

func (c *blazarSelect) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarSelect {
	c.MixOnDoubleClick(function)
	return c
}

func (c *blazarSelect) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarSelect {
	c.MixOnMouseUp(function)
	return c
}

func (c *blazarSelect) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarSelect {
	c.MixOnMouseDown(function)
	return c
}

func (c *blazarSelect) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarSelect {
	c.MixOnKeyDown(function)
	return c
}

func (c *blazarSelect) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarSelect {
	c.MixOnKeyPress(function)
	return c
}

func (c *blazarSelect) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarSelect {
	c.MixOnKeyUp(function)
	return c
}
