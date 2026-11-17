package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// On invokes the specified handler when the corresponding event is triggered.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCopy) On(name string, eventHandler app.EventHandler, options ...app.EventOption) *blazarCopy {
	c.MixOn(name, eventHandler, options...)
	return c
}

// OnClick invokes the specified handler when the element is clicked.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarCopy) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarCopy {
	c.MixOnClick(function)
	return c
}

// OnDoubleClick invokes the specified handler when the element is double-clicked.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarCopy) OnDoubleClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarCopy {
	c.MixOnDoubleClick(function)
	return c
}

// OnMouseUp invokes the specified handler when the mouse button is released.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarCopy) OnMouseUp(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarCopy {
	c.MixOnMouseUp(function)
	return c
}

// OnMouseDown invokes the specified handler when the mouse button is pressed.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarCopy) OnMouseDown(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarCopy {
	c.MixOnMouseDown(function)
	return c
}

// OnKeyDown invokes the specified handler when a key is pressed.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarCopy) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarCopy {
	c.MixOnKeyDown(function)
	return c
}

// OnKeyPress invokes the specified handler when a key is pressed.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarCopy) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarCopy {
	c.MixOnKeyPress(function)
	return c
}

// OnKeyUp invokes the specified handler when a key is released.
//
// This uses the `htmlevent` package to simplify the event handling.
func (c *blazarCopy) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarCopy {
	c.MixOnKeyUp(function)
	return c
}
