package blazar

import (
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Form() *blazarForm {
	return &blazarForm{
		ISpacer: true,
	}
}

type blazarForm struct {
	app.Compo
	UseEvents

	IClasses []string
	IStyles  map[string]string

	ISpacer         bool
	IBody           []app.UI
	ICancelFunction func(ctx app.Context)
	ICancelLabel    string
	ICancelIcon     string
	ISubmitFunction func(ctx app.Context)
	ISubmitLabel    string
	ISubmitIcon     string
	IActions        []FormAction
}

type FormAction struct {
	Name     string
	Icon     string
	To       string
	Function func(ctx app.Context)
}

var _ app.Composer = (*blazarForm)(nil)

func (c *blazarForm) Class(class ...string) *blazarForm {
	c.IClasses = class
	return c
}

func (c *blazarForm) Spacer(spacer bool) *blazarForm {
	c.ISpacer = spacer
	return c
}

func (c *blazarForm) Style(name, value string) *blazarForm {
	if c.IStyles == nil {
		c.IStyles = make(map[string]string)
	}
	c.IStyles[name] = value
	return c
}

func (c *blazarForm) Action(actions ...FormAction) *blazarForm {
	c.IActions = actions
	return c
}

func (c *blazarForm) CancelFunction(function func(ctx app.Context)) *blazarForm {
	c.ICancelFunction = function
	return c
}

func (c *blazarForm) CancelLabel(label string) *blazarForm {
	c.ICancelLabel = label
	return c
}

func (c *blazarForm) CancelIcon(icon string) *blazarForm {
	c.ICancelIcon = icon
	return c
}

func (c *blazarForm) SubmitFunction(function func(ctx app.Context)) *blazarForm {
	c.ISubmitFunction = function
	return c
}

func (c *blazarForm) SubmitIcon(icon string) *blazarForm {
	c.ISubmitIcon = icon
	return c
}

func (c *blazarForm) SubmitLabel(label string) *blazarForm {
	c.ISubmitLabel = label
	return c
}

func (c *blazarForm) Body(body ...app.UI) *blazarForm {
	c.IBody = body
	return c
}

func (c *blazarForm) On(event string, function func(ctx app.Context, e app.Event)) *blazarForm {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarForm) Render() app.UI {
	element := app.Div().
		Class(append([]string{"blazar-form"}, c.IClasses...)...).
		Body(
			c.UseEvents.Wrap(
				app.Div().
					Class("blazar-form__form").
					On("keyup", func(ctx app.Context, e app.Event) {
						ctx.PreventUpdate()

						slog.InfoContext(ctx.Context, "blazarForm: Keypress", "key", e.Get("key").String())

						// If the user pressed "Enter", then perform the default action.
						//
						// If set, the default action is the submit function.
						// Otherwise, the default action is the *last* custom action.
						switch e.Get("key").String() {
						case "Enter":
							if c.ISubmitFunction != nil {
								c.ISubmitFunction(ctx)
							} else {
								if len(c.IActions) > 0 {
									lastAction := c.IActions[len(c.IActions)-1]
									if lastAction.Function != nil {
										lastAction.Function(ctx)
									} else if lastAction.To != "" {
										ctx.Navigate(lastAction.To)
									}
								}
							}
						case "Escape":
							if c.ICancelFunction != nil {
								c.ICancelFunction(ctx)
							}
						}
					}).
					Body(
						c.IBody...,
					),
			),
			app.Div().
				Class("blazar-form__actions").
				Body(
					app.If(c.ICancelFunction != nil, func() app.UI {
						return Button().
							Flat(true).
							Label(func() string {
								if c.ICancelLabel != "" {
									return c.ICancelLabel
								}
								return "Cancel"
							}()).
							Icon(c.ICancelIcon).
							On("click", func(ctx app.Context, e app.Event) {
								c.ICancelFunction(ctx)
							})
					}),
					app.If(c.ISpacer, func() app.UI {
						return app.Span().Style("flex", "1")
					}),
					app.Range(c.IActions).Slice(func(i int) app.UI {
						action := c.IActions[i]
						return Button().
							Flat(false).
							Label(action.Name).
							Icon(action.Icon).
							To(action.To).
							On("click", func(ctx app.Context, e app.Event) {
								if action.Function == nil {
									ctx.PreventUpdate()
									return
								}
								action.Function(ctx)
							})
					}),
					app.If(c.ISubmitFunction != nil, func() app.UI {
						return Button().
							Flat(false).
							Label(func() string {
								if c.ISubmitLabel != "" {
									return c.ISubmitLabel
								}
								return "Submit"
							}()).
							Icon(c.ISubmitIcon).
							On("click", func(ctx app.Context, e app.Event) {
								c.ISubmitFunction(ctx)
							})
					}),
				),
		)
	for name, value := range c.IStyles {
		element = element.Style(name, value)
	}
	return element
}
