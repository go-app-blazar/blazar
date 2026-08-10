package blazar

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// Copy creates a new copy component.
// The copy component is a button that allows the user to copy text to the clipboard.
func Copy() *blazarCopy {
	return &blazarCopy{
		IIcon:   "copy",
		IStyles: map[string]string{},
	}
}

type blazarCopy struct {
	app.Compo
	IDisabled bool                                                // If true, copying will be disabled.
	IIcon     string                                              // The icon to display next to the content.
	IText     string                                              // The text to display.
	IBody     []app.UI                                            // The body to display inside the copy component (optional).
	IStyles   map[string]string                                   // The styles to apply to the copy component.
	IValue    string                                              // This is the value that will be copied to the clipboard.
	IOnCopy   func(ctx app.Context, value string)                 // The function to call when the text has been copied.
	IOnClick  func(ctx app.Context, event htmlevent.PointerEvent) // The function to call when the copy is clicked.
}

var _ app.Composer = (*blazarCopy)(nil)

func (c *blazarCopy) Disabled(disabled bool) *blazarCopy {
	c.IDisabled = disabled
	return c
}

func (c *blazarCopy) Icon(icon string) *blazarCopy {
	c.IIcon = icon
	return c
}

func (c *blazarCopy) Style(name, value string) *blazarCopy {
	c.IStyles[name] = value
	return c
}

func (c *blazarCopy) Text(text string) *blazarCopy {
	c.IText = text
	return c
}

func (c *blazarCopy) Value(value string) *blazarCopy {
	c.IValue = value
	return c
}

func (c *blazarCopy) Body(body ...app.UI) *blazarCopy {
	c.IBody = body
	return c
}

// OnClick is called when the copy is clicked.
func (c *blazarCopy) OnClick(function func(ctx app.Context, event htmlevent.PointerEvent)) *blazarCopy {
	c.IOnClick = function
	return c
}

// OnCopy is called when the text has been copied.
// The text value is passed to the function.
func (c *blazarCopy) OnCopy(function func(ctx app.Context, value string)) *blazarCopy {
	c.IOnCopy = function
	return c
}

func (c *blazarCopy) Render() app.UI {
	if debugCopy {
		slog.DebugContext(context.TODO(), "blazarCopy: Render", "value", c.IValue)
	}

	disabledClass := ""
	if c.IDisabled {
		disabledClass = "disabled"
	}

	element := app.Div().
		Class("blazar-copy", disabledClass).
		On("click", func(ctx app.Context, e app.Event) {
			if debugCopy {
				slog.DebugContext(ctx.Context, "Collapse: OnClick", "disabled", c.IDisabled, "value", c.IValue)
			}
			if c.IDisabled {
				return
			}
			if c.IOnClick != nil {
				var pointerEvent htmlevent.PointerEvent
				htmlevent.MustParse(e, &pointerEvent)
				c.IOnClick(ctx, pointerEvent)
			}
			app.Window().Get("navigator").Get("clipboard").Call("writeText", c.IValue).Then(func(result app.Value) {
				if c.IOnCopy != nil {
					c.IOnCopy(ctx, c.IValue)
				}
			})
			ctx.Update()
		}).
		Body(
			app.Div().
				Class("blazar-copy__wrapper").
				Body(
					app.If(c.IText != "", func() app.UI {
						return app.Div().
							Class("blazar-copy__content").
							Text(c.IText)
					}).Else(func() app.UI {
						return app.Div().
							Class("blazar-copy__content").
							Body(c.IBody...)
					}),
					app.If(c.IIcon != "", func() app.UI {
						return Button().
							Class("blazar-copy__icon").
							Icon(c.IIcon).
							Flat(true).
							Disabled(c.IDisabled)
					}),
				),
		)
	for name, value := range c.IStyles {
		element = element.Style(name, value)
	}
	return element
}
