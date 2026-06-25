package blazar

import (
	"context"
	"log/slog"
	"slices"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Multiselect() *blazarMultiselect {
	return &blazarMultiselect{}
}

type blazarMultiselect struct {
	app.Compo
	UseEvents
	IName              string
	ILabel             string
	IDisabled          bool
	IAllowedValues     []SelectOption
	ISelectedValues    []string
	BindSelectedValues *[]string
}

var _ app.Composer = (*blazarMultiselect)(nil)
var _ app.Updater = (*blazarMultiselect)(nil)

func (c *blazarMultiselect) OnUpdate(ctx app.Context) {
	if debugMultiselect {
		slog.DebugContext(ctx.Context, "blazarMultiselect: OnUpdate")
	}
}

func (c *blazarMultiselect) Disabled(disabled bool) *blazarMultiselect {
	c.IDisabled = disabled
	return c
}

func (c *blazarMultiselect) Name(name string) *blazarMultiselect {
	c.IName = name
	return c
}

func (c *blazarMultiselect) Label(label string) *blazarMultiselect {
	c.ILabel = label
	return c
}

func (c *blazarMultiselect) AllowedValue(allowedValue ...SelectOption) *blazarMultiselect {
	c.IAllowedValues = allowedValue
	return c
}

func (c *blazarMultiselect) SelectedValue(selectedValue ...string) *blazarMultiselect {
	c.ISelectedValues = selectedValue
	if c.BindSelectedValues != nil {
		*c.BindSelectedValues = selectedValue
	}
	return c
}

func (c *blazarMultiselect) Bind(bindSelectedValues *[]string) *blazarMultiselect {
	c.BindSelectedValues = bindSelectedValues
	if c.BindSelectedValues != nil {
		c.ISelectedValues = make([]string, len(*c.BindSelectedValues))
		copy(c.ISelectedValues, *c.BindSelectedValues)
	}
	return c
}

func (c *blazarMultiselect) On(event string, function func(ctx app.Context, e app.Event)) *blazarMultiselect {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarMultiselect) Render() app.UI {
	if debugMultiselect {
		slog.DebugContext(context.TODO(), "blazarMultiselect: Render", "label", c.ILabel, "allowedValues", c.IAllowedValues, "selectedValues", c.ISelectedValues, "bindSelectedValues", c.BindSelectedValues)
	}
	return InputWrapper().
		Class("blazar-multiselect").
		Label(c.ILabel).
		Body(
			c.UseEvents.Wrap(
				app.Select().
					Name(c.IName).
					Class("blazar-multiselect__select").
					Disabled(c.IDisabled).
					Multiple(true).
					Body(
						app.Range(c.IAllowedValues).Slice(func(i int) app.UI {
							allowedValue := c.IAllowedValues[i]
							return app.Option().
								Value(allowedValue.Value).
								Text(allowedValue.Label).
								Disabled(allowedValue.Disabled).
								Selected(slices.Contains(c.ISelectedValues, allowedValue.Value))
						}),
					),
				WithOn("change", func(ctx app.Context, e app.Event) {
					SelectedValuesTo(&c.ISelectedValues)(ctx, e)
					if c.BindSelectedValues != nil {
						*c.BindSelectedValues = make([]string, len(c.ISelectedValues))
						copy(*c.BindSelectedValues, c.ISelectedValues)
					}
					ctx.PreventUpdate()
				}),
			),
		)
}

// SelectedValuesTo is an event handler that updates the variable based on the selected options in a multiselect.
func SelectedValuesTo(selectedValues *[]string) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		targetElement := e.Value.Get("target")
		if targetElement.IsNull() {
			return
		}
		selectedOptions := targetElement.Get("selectedOptions")
		if selectedOptions.IsNull() {
			return
		}

		*selectedValues = []string{}
		selectedOptionsLength := selectedOptions.Length()
		for i := range selectedOptionsLength {
			selectedOption := selectedOptions.Index(i)
			if selectedOption.IsNull() {
				continue
			}
			selectedOptionValue := selectedOption.Get("value").String()
			*selectedValues = append(*selectedValues, selectedOptionValue)
		}
		if debugMultiselect {
			slog.DebugContext(ctx.Context, "SelectedValuesTo: Selected values", "selectedValues", *selectedValues)
		}
	}
}
