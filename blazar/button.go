package blazar

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Button() *blazarButton {
	return &blazarButton{}
}

type blazarButton struct {
	app.Compo
	UseEvents
	IFlat     bool
	IIcon     string
	ILabel    string
	ITo       string
	IRound    bool
	IDisabled bool
}

var _ app.Composer = (*blazarButton)(nil)

func (c *blazarButton) Disabled(disabled bool) *blazarButton {
	c.IDisabled = disabled
	return c
}

func (c *blazarButton) Flat(flat bool) *blazarButton {
	c.IFlat = flat
	return c
}

func (c *blazarButton) Icon(icon string) *blazarButton {
	c.IIcon = icon
	return c
}

func (c *blazarButton) Label(label string) *blazarButton {
	c.ILabel = label
	return c
}

func (c *blazarButton) Round(round bool) *blazarButton {
	c.IRound = round
	return c
}

func (c *blazarButton) To(to string) *blazarButton {
	c.ITo = to
	return c
}

func (c *blazarButton) On(event string, function func(ctx app.Context, e app.Event)) *blazarButton {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarButton) Render() app.UI {
	disabledClass := ""
	if c.IDisabled {
		disabledClass = "disabled"
	}

	flatClass := ""
	if c.IFlat {
		flatClass = "flat"
	}

	roundClass := ""
	if c.IRound {
		roundClass = "round"
	}

	var body []app.UI
	if c.IIcon != "" {
		body = append(body, Icon().
			Class("blazar-button__icon").
			Icon(c.IIcon),
		)
	}
	if c.ILabel != "" {
		body = append(body, app.Span().
			Class("blazar-button__label").
			Body(
				app.Span().
					Text(c.ILabel),
			),
		)
	}

	var innerElement app.UI
	if c.ITo == "" || c.IDisabled {
		innerElement = app.Span().
			Class("blazar-button__content").
			Body(body...)
	} else {
		innerElement = app.A().
			Class("blazar-button__content").
			Href(c.ITo).
			Body(body...)
	}

	var element app.UI
	element = app.Span().
		Class("blazar-button", disabledClass, roundClass, flatClass).
		TabIndex(0).
		Role("button").
		Body(innerElement).
		OnKeyPress(func(ctx app.Context, e app.Event) {
			if e.Get("key").String() == "Enter" || e.Get("key").String() == " " {
				e.Get("target").Call("click")
			}
		})
	if !c.IDisabled {
		element = c.UseEvents.Wrap(element)
	}
	return element
}
