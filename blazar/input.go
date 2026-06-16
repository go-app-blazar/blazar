package blazar

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Input[T any]() *blazarInput[T] {
	return &blazarInput[T]{}
}

type blazarInput[T any] struct {
	app.Compo
	UseEvents
	IAutoFocus   bool
	IType        string
	IName        string
	IDisabled    bool
	ILabel       string
	IPlaceholder string
	IValue       T
	BindValue    *T
}

var _ app.Composer = (*blazarInput[any])(nil)

func (c *blazarInput[T]) AutoFocus(autoFocus bool) *blazarInput[T] {
	c.IAutoFocus = autoFocus
	return c
}

func (c *blazarInput[T]) Name(name string) *blazarInput[T] {
	c.IName = name
	return c
}

func (c *blazarInput[T]) Placeholder(placeholder string) *blazarInput[T] {
	c.IPlaceholder = placeholder
	return c
}

func (c *blazarInput[T]) Disabled(disabled bool) *blazarInput[T] {
	c.IDisabled = disabled
	return c
}

func (c *blazarInput[T]) Type(inputType string) *blazarInput[T] {
	c.IType = inputType
	return c
}

func (c *blazarInput[T]) Label(label string) *blazarInput[T] {
	c.ILabel = label
	return c
}

func (c *blazarInput[T]) Value(value T) *blazarInput[T] {
	if c.BindValue == nil {
		c.BindValue = new(T)
	}
	*c.BindValue = value
	return c
}

func (c *blazarInput[T]) Bind(valuePointer *T) *blazarInput[T] {
	c.IValue = *valuePointer
	c.BindValue = valuePointer
	return c
}

func (c *blazarInput[T]) On(event string, function func(ctx app.Context, e app.Event)) *blazarInput[T] {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarInput[T]) Render() app.UI {
	slog.InfoContext(context.TODO(), "blazarInput: Render", "label", c.ILabel, "type", c.IType, "value", c.BindValue, "placeholder", c.IPlaceholder, "disabled", c.IDisabled)

	kind := reflect.TypeOf(c.IValue).Kind()

	var minValue any
	inputType := "text"
	{
		switch kind {
		case reflect.Bool:
			inputType = "checkbox"
		case reflect.Float32:
			inputType = "number"
		case reflect.Float64:
			inputType = "number"
		case reflect.String:
			inputType = "text"
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			inputType = "number"
			minValue = 0
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			inputType = "number"
		}
	}

	if c.IType != "" {
		inputType = c.IType
	}

	var checked bool
	var value any
	if inputType == "checkbox" {
		checked = fmt.Sprintf("%v", c.IValue) == "true"
	} else {
		value = fmt.Sprintf("%v", c.IValue)
		if c.BindValue != nil {
			value = fmt.Sprintf("%v", *c.BindValue)
		}
	}

	return InputWrapper().
		Class("blazar-input").
		Label(c.ILabel).
		Body(
			c.UseEvents.Wrap(
				app.Input().
					Class("blazar-input__input").
					Disabled(c.IDisabled).
					ReadOnly(c.IDisabled).
					AutoComplete(false).
					AutoFocus(c.IAutoFocus).
					Name(c.IName).
					Type(inputType).
					Checked(checked).
					Value(value).
					Min(minValue).
					Placeholder(c.IPlaceholder),
				WithOn("change", func(ctx app.Context, e app.Event) {
					//slog.InfoContext(ctx.Context, "blazarInput: Change", "value", value)
					//slog.InfoContext(ctx.Context, "blazarInput: Change", "e.target.checked", e.Get("target").Get("checked").String())

					if c.BindValue != nil {
						if kind == reflect.Bool {
							boolValue := reflect.ValueOf(e.Get("target").Get("checked").Bool())
							*c.BindValue = boolValue.Convert(reflect.TypeOf(c.IValue)).Interface().(T)
						} else {
							c.ValueTo(c.BindValue)(ctx, e)
						}
					}
				}),
				WithOn("keypress", func(ctx app.Context, e app.Event) {
					ctx.PreventUpdate()

					if e.Get("key").String() == "Enter" {
						if c.BindValue != nil {
							if kind == reflect.Bool {
								boolValue := reflect.ValueOf(e.Get("target").Get("checked").Bool())
								*c.BindValue = boolValue.Convert(reflect.TypeOf(c.IValue)).Interface().(T)
							} else {
								c.ValueTo(c.BindValue)(ctx, e)
							}
						}
					}
				}),
			),
		)
}
