package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// On invokes the specified handler when the corresponding event is triggered.
//
// This is for compatibility with go-app HTML components.
func (c *blazarSelect) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarSelect {
	c.MixOn(name, eventHandler, options...)
	return c
}

// OnClick invokes the specified handler when the element is clicked.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarSelect) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarSelect {
	c.MixOnClick(function)
	return c
}

// OnDoubleClick invokes the specified handler when the element is double-clicked.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarSelect) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarSelect {
	c.MixOnDoubleClick(function)
	return c
}

// OnMouseUp invokes the specified handler when the mouse button is released.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarSelect) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarSelect {
	c.MixOnMouseUp(function)
	return c
}

// OnMouseDown invokes the specified handler when the mouse button is pressed.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarSelect) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarSelect {
	c.MixOnMouseDown(function)
	return c
}

// OnKeyDown invokes the specified handler when a key is pressed.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarSelect) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarSelect {
	c.MixOnKeyDown(function)
	return c
}

// OnKeyPress invokes the specified handler when a key is pressed.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarSelect) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarSelect {
	c.MixOnKeyPress(function)
	return c
}

// OnKeyUp invokes the specified handler when a key is released.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarSelect) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarSelect {
	c.MixOnKeyUp(function)
	return c
}
