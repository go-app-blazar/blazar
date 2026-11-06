package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Button() *blazarButton {
	return &blazarButton{}
}

type blazarButton struct {
	app.Compo
	UseEvents
	MixHTML
	IAutoFocus bool
	IFlat      bool
	IOutline   bool
	IIcon      string
	ILabel     string
	ITo        string
	ITarget    string // When "To" is set, this will be the target of the button.
	IRound     bool
	IDisabled  bool
}

var _ app.Composer = (*blazarButton)(nil)
var _ app.Mounter = (*blazarButton)(nil)

func (c *blazarButton) OnMount(ctx app.Context) {
	if c.IAutoFocus {
		c.JSValue().Call("focus")
	}
}

// BEGIN: mixHTML methods

func (c *blazarButton) Attr(name string, value any) *blazarButton {
	c.mixAttr(name, value)
	return c
}

func (c *blazarButton) Class(class ...string) *blazarButton {
	c.mixClass(class...)
	return c
}

func (c *blazarButton) DataSet(name string, value any) *blazarButton {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarButton) DataSets(dataSets map[string]any) *blazarButton {
	c.mixDataSets(dataSets)
	return c
}

func (c *blazarButton) Style(name, value string) *blazarButton {
	c.mixStyle(name, value)
	return c
}

func (c *blazarButton) Styles(styles map[string]string) *blazarButton {
	c.mixStyles(styles)
	return c
}

func (c *blazarButton) Title(title string) *blazarButton {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

func (c *blazarButton) AutoFocus(autoFocus bool) *blazarButton {
	c.IAutoFocus = autoFocus
	return c
}

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

func (c *blazarButton) Outline(outline bool) *blazarButton {
	c.IOutline = outline
	return c
}

func (c *blazarButton) Round(round bool) *blazarButton {
	c.IRound = round
	return c
}

func (c *blazarButton) Target(target string) *blazarButton {
	c.ITarget = target
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

func (c *blazarButton) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarButton {
	c.UseEvents.OnClick(function)
	return c
}

func (c *blazarButton) OnKeyPress(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarButton {
	c.UseEvents.OnKeyPress(function)
	return c
}

func (c *blazarButton) OnKeyDown(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarButton {
	c.UseEvents.OnKeyDown(function)
	return c
}
func (c *blazarButton) OnKeyUp(function func(ctx app.Context, e htmlevent.KeyboardEvent)) *blazarButton {
	c.UseEvents.OnKeyUp(function)
	return c
}

func (c *blazarButton) Render() app.UI {
	disabledClass := ""
	if c.IDisabled {
		disabledClass = "blazar-button--disabled"
	}

	flatClass := ""
	if c.IFlat {
		flatClass = "blazar-button--flat"
	}

	outlineClass := ""
	if c.IOutline {
		outlineClass = "blazar-button--outline"
	}

	roundClass := ""
	if c.IRound {
		roundClass = "blazar-button--round"
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
			Target(c.ITarget).
			Body(body...)
	}

	classes := []string{"blazar-button", disabledClass, roundClass, flatClass, outlineClass}

	var element app.HTMLSpan
	element = app.Span().
		Class(classes...).
		TabIndex(0).
		Role("button").
		Body(innerElement).
		Title(c.ITitle).
		OnKeyPress(htmlevent.WrapKeyboardEvent(func(ctx app.Context, e htmlevent.KeyboardEvent) {
			if e.Key == "Enter" || e.Key == " " {
				e.Target.Call("click")

				e.PreventDefault()
			}
		}))
	if c.IAutoFocus {
		element = element.Attr("autofocus", "true")
	}

	var finalElement app.UI
	finalElement = c.MixHTML.Wrap(element)
	if !c.IDisabled {
		finalElement = c.UseEvents.Wrap(finalElement)
	}
	return finalElement
}
