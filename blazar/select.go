package blazar

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Select() *blazarSelect {
	return &blazarSelect{}
}

type blazarSelect struct {
	app.Compo
	UseEvents
	IName             string
	ILabel            string
	IAllowedValues    []SelectOption
	ISelectedValue    string
	BindSelectedValue *string
}

var _ app.Composer = (*blazarSelect)(nil)
var _ app.Updater = (*blazarSelect)(nil)

func (c *blazarSelect) OnUpdate(ctx app.Context) {
	slog.InfoContext(ctx.Context, "blazarSelect: OnUpdate")
}

func (c *blazarSelect) Name(name string) *blazarSelect {
	c.IName = name
	return c
}

func (c *blazarSelect) Label(label string) *blazarSelect {
	c.ILabel = label
	return c
}

func (c *blazarSelect) AllowedValue(allowedValue ...SelectOption) *blazarSelect {
	c.IAllowedValues = allowedValue
	return c
}

func (c *blazarSelect) SelectedValue(selectedValue string) *blazarSelect {
	c.ISelectedValue = selectedValue
	if c.BindSelectedValue != nil {
		*c.BindSelectedValue = selectedValue
	}
	return c
}

func (c *blazarSelect) Bind(bindSelectedValue *string) *blazarSelect {
	c.BindSelectedValue = bindSelectedValue
	if c.BindSelectedValue != nil {
		c.ISelectedValue = *c.BindSelectedValue
	}
	return c
}

func (c *blazarSelect) On(event string, function func(ctx app.Context, e app.Event)) *blazarSelect {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarSelect) Render() app.UI {
	slog.InfoContext(context.TODO(), "blazarSelect: Render", "label", c.ILabel, "allowedValues", c.IAllowedValues, "selectedValue", c.ISelectedValue, "bindSelectedValue", c.BindSelectedValue)
	return InputWrapper().
		Class("blazar-select").
		Label(c.ILabel).
		Body(
			c.UseEvents.Wrap(
				app.Select().
					Name(c.IName).
					Class("blazar-select__select").
					Multiple(false).
					Body(
						app.Range(c.IAllowedValues).Slice(func(i int) app.UI {
							allowedValue := c.IAllowedValues[i]
							return app.Option().
								Value(allowedValue.Value).
								Text(allowedValue.Label).
								Disabled(allowedValue.Disabled).
								Selected(c.ISelectedValue == allowedValue.Value)
						}),
					),
				WithOn("change", func(ctx app.Context, e app.Event) {
					c.ValueTo(&c.ISelectedValue)(ctx, e)
					if c.BindSelectedValue != nil {
						*c.BindSelectedValue = c.ISelectedValue
					}
					ctx.PreventUpdate()
				}),
			),
		)
}
