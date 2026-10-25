package blazar

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Checkbox() *blazarCheckbox {
	return &blazarCheckbox{}
}

type blazarCheckbox struct {
	app.Compo
	mixHTML

	IOnChange  func(ctx app.Context, checked bool)
	IAutoFocus bool
	IDisabled  bool
	ILabel     string
	IValue     bool
	bindValue  *bool
}

var _ app.Composer = (*blazarCheckbox)(nil)

// BEGIN: mixHTML methods

func (c *blazarCheckbox) Attr(name string, value any) *blazarCheckbox {
	c.mixAttr(name, value)
	return c
}

func (c *blazarCheckbox) Class(class ...string) *blazarCheckbox {
	c.mixClass(class...)
	return c
}

func (c *blazarCheckbox) DataSet(name string, value any) *blazarCheckbox {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarCheckbox) Style(name, value string) *blazarCheckbox {
	c.mixStyle(name, value)
	return c
}

func (c *blazarCheckbox) Title(title string) *blazarCheckbox {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

func (c *blazarCheckbox) AutoFocus(autoFocus bool) *blazarCheckbox {
	c.IAutoFocus = autoFocus
	return c
}

func (c *blazarCheckbox) Disabled(disabled bool) *blazarCheckbox {
	c.IDisabled = disabled
	return c
}

func (c *blazarCheckbox) Label(label string) *blazarCheckbox {
	c.ILabel = label
	return c
}

func (c *blazarCheckbox) Value(value bool) *blazarCheckbox {
	if c.bindValue == nil {
		c.bindValue = new(bool)
	}
	*c.bindValue = value
	return c
}

func (c *blazarCheckbox) Bind(valuePointer *bool) *blazarCheckbox {
	c.IValue = *valuePointer
	c.bindValue = valuePointer
	return c
}

func (c *blazarCheckbox) OnChange(function func(ctx app.Context, checked bool)) *blazarCheckbox {
	c.IOnChange = function
	return c
}

func (c *blazarCheckbox) Render() app.UI {
	if debugCheckbox {
		slog.DebugContext(context.TODO(), "blazarCheckbox: Render", "label", c.ILabel, "value", c.IValue, "bindValue", c.bindValue, "disabled", c.IDisabled)
	}

	checked := c.IValue
	if c.bindValue != nil {
		checked = *c.bindValue
	}

	checkedClass := ""
	if checked {
		checkedClass = "checked"
	}

	disabledClass := ""
	if c.IDisabled {
		disabledClass = "blazar-checkbox--disabled"
	}

	iconName := "square"
	if checked {
		iconName = "square-check"
	}

	// TODO: How do we handle autofocus?
	// TDOO: How do we handle "name" for forms?

	icon := Icon().
		Class("blazar-checkbox__icon").
		Icon(iconName).
		Solid(checked)
	if !c.IDisabled {
		icon = icon.On("click", func(ctx app.Context, e app.Event) {
			if c.bindValue != nil {
				*c.bindValue = !*c.bindValue
				if debugCheckbox {
					slog.DebugContext(ctx.Context, "blazarCheckbox: Click", "old bindValue", *c.bindValue)
				}

				if c.IOnChange != nil {
					c.IOnChange(ctx, *c.bindValue)
				}
			}
		})
		icon = icon.On("keyup", func(ctx app.Context, e app.Event) {
			ctx.PreventUpdate()

			if e.Get("key").String() == "Enter" {
				if c.bindValue != nil {
					*c.bindValue = !*c.bindValue
					if debugCheckbox {
						slog.DebugContext(ctx.Context, "blazarCheckbox: Keyup", "old bindValue", *c.bindValue)
					}

					if c.IOnChange != nil {
						c.IOnChange(ctx, *c.bindValue)
					}
				}
			}
		})
	}

	var finalElement app.UI
	finalElement = InputWrapper().
		Class("blazar-checkbox", checkedClass, disabledClass).
		Label(c.ILabel).
		Body(
			icon,
		)

	finalElement = c.mixHTML.Wrap(finalElement)

	return finalElement
}
