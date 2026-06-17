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
	IOnChange  func(ctx app.Context, checked bool)
	IAutoFocus bool
	IDisabled  bool
	ILabel     string
	IValue     bool
	bindValue  *bool
}

var _ app.Composer = (*blazarCheckbox)(nil)

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
	slog.InfoContext(context.TODO(), "blazarCheckbox: Render", "label", c.ILabel, "value", c.IValue, "bindValue", c.bindValue, "disabled", c.IDisabled)

	disabledClass := ""
	if c.IDisabled {
		disabledClass = "disabled"
	}

	checked := c.IValue
	if c.bindValue != nil {
		checked = *c.bindValue
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
				slog.InfoContext(ctx.Context, "blazarCheckbox: Click", "old bindValue", *c.bindValue)

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
					slog.InfoContext(ctx.Context, "blazarCheckbox: Keyup", "old bindValue", *c.bindValue)

					if c.IOnChange != nil {
						c.IOnChange(ctx, *c.bindValue)
					}
				}
			}
		})
	}

	return InputWrapper().
		Class("blazar-checkbox", disabledClass).
		Label(c.ILabel).
		Body(
			icon,
		)
}
