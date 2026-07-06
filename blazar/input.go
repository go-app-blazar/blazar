package blazar

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"strconv"
	"strings"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Input[T any]() *blazarInput[T] {
	return &blazarInput[T]{
		IOutline: true,
	}
}

type blazarInput[T any] struct {
	app.Compo
	UseEvents
	UseData
	IAutoFocus   bool
	IClearable   bool
	IOutline     bool
	IPrefix      string
	ISuffix      string
	IType        string
	IName        string
	IDisabled    bool
	ILabel       string
	IPlaceholder string
	IValue       T
	IMinValue    *T
	IMaxValue    *T
	IStepValue   *T
	BindValue    *T
}

var _ app.Composer = (*blazarInput[any])(nil)

func (c *blazarInput[T]) AutoFocus(autoFocus bool) *blazarInput[T] {
	c.IAutoFocus = autoFocus
	return c
}

func (c *blazarInput[T]) Clearable(clearable bool) *blazarInput[T] {
	c.IClearable = clearable
	return c
}

func (c *blazarInput[T]) Name(name string) *blazarInput[T] {
	c.IName = name
	return c
}

func (c *blazarInput[T]) Outline(outline bool) *blazarInput[T] {
	c.IOutline = outline
	return c
}

func (c *blazarInput[T]) Placeholder(placeholder string) *blazarInput[T] {
	c.IPlaceholder = placeholder
	return c
}

func (c *blazarInput[T]) Prefix(prefix string) *blazarInput[T] {
	c.IPrefix = prefix
	return c
}

func (c *blazarInput[T]) Suffix(suffix string) *blazarInput[T] {
	c.ISuffix = suffix
	return c
}

func (c *blazarInput[T]) Min(minValue T) *blazarInput[T] {
	c.IMinValue = &minValue
	return c
}

func (c *blazarInput[T]) Max(maxValue T) *blazarInput[T] {
	c.IMaxValue = &maxValue
	return c
}

func (c *blazarInput[T]) Step(stepValue T) *blazarInput[T] {
	c.IStepValue = &stepValue
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

func (c *blazarInput[T]) DataSet(name string, value any) *blazarInput[T] {
	c.UseData.DataSet(name, value)
	return c
}

func (c *blazarInput[T]) On(event string, function func(ctx app.Context, e app.Event)) *blazarInput[T] {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarInput[T]) Render() app.UI {
	derefOrNil := func(value *T) string {
		if value == nil {
			return "nil"
		}
		return fmt.Sprintf("%v", *value)
	}
	if debugInput {
		slog.DebugContext(context.TODO(), "blazarInput: Render", "label", c.ILabel, "type", c.IType, "value", c.IValue, "bindValue", derefOrNil(c.BindValue), "placeholder", c.IPlaceholder, "disabled", c.IDisabled)
	}

	kind := reflect.TypeOf(c.IValue).Kind()

	var minValue any
	var maxValue any
	var stepValue float64
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
	if c.IMinValue != nil {
		minValue = *c.IMinValue
	}
	if c.IMaxValue != nil {
		maxValue = *c.IMaxValue
	}
	if c.IStepValue != nil {
		v := *c.IStepValue
		vString := ""
		if kind == reflect.Float32 || kind == reflect.Float64 {
			vString = fmt.Sprintf("%f", reflect.ValueOf(v).Float())
		} else {
			vString = fmt.Sprintf("%v", v)
		}
		floatValue, err := strconv.ParseFloat(vString, 64)
		if err != nil {
			// Oh well.
		} else {
			stepValue = floatValue
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
		if kind == reflect.Float32 || kind == reflect.Float64 {
			stringValue := fmt.Sprintf("%f", reflect.ValueOf(c.IValue).Float())
			if c.BindValue != nil {
				stringValue = fmt.Sprintf("%f", reflect.ValueOf(*c.BindValue).Float())
			}
			if strings.Contains(stringValue, ".") {
				stringValue = strings.TrimRight(stringValue, "0")
			}
			if strings.HasPrefix(stringValue, ".") {
				stringValue = "0" + stringValue
			}
			if strings.HasSuffix(stringValue, ".") {
				stringValue = stringValue + "0"
			}
			if strings.HasSuffix(stringValue, ".0") {
				stringValue = strings.TrimSuffix(stringValue, ".0")
			}
			value = stringValue
		} else {
			value = fmt.Sprintf("%v", c.IValue)
			if c.BindValue != nil {
				value = fmt.Sprintf("%v", *c.BindValue)
			}
		}
	}

	disabledClass := ""
	if c.IDisabled {
		disabledClass = "blazar-input-wrapper--disabled"
	}

	outlineClass := ""
	if c.IOutline {
		outlineClass = "blazar-input-wrapper--outline"
	}

	isZeroValue := reflect.ValueOf(c.IValue).IsZero()

	return InputWrapper().
		Class("blazar-input", disabledClass, outlineClass).
		Label(c.ILabel).
		Body(
			app.If(c.IPrefix != "", func() app.UI {
				return app.Span().Class("blazar-input__prefix").Text(c.IPrefix)
			}),
			c.UseEvents.Wrap(
				c.UseData.Wrap(
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
						Max(maxValue).
						Step(stepValue).
						Placeholder(c.IPlaceholder),
				),
				WithOn("change", func(ctx app.Context, e app.Event) {
					/*
						if debugInput {
							slog.DebugContext(ctx.Context, "blazarInput: Change", "value", value)
							slog.DebugContext(ctx.Context, "blazarInput: Change", "e.target.checked", e.Get("target").Get("checked").String())
						}
					*/

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
			app.If(c.ISuffix != "", func() app.UI {
				return app.Span().Class("blazar-input__suffix").Text(c.ISuffix)
			}),
			app.If(c.IClearable && !isZeroValue, func() app.UI {
				return Button().
					Class("blazar-input__clear-icon").
					Icon("circle-xmark").
					Flat(true).
					Disabled(c.IDisabled).
					On("click", func(ctx app.Context, e app.Event) {
						if c.BindValue != nil {
							*c.BindValue = reflect.Zero(reflect.TypeOf(*c.BindValue)).Interface().(T)
						}
					})
			}),
		)
}
