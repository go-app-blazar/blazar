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

	loading bool
}

type FormAction struct {
	Name            string                // The name of the action.
	Icon            string                // The icon of the action.  If empty, then no icon will be shown.
	To              string                // If set, then the action will navigate to the target URL.
	Target          string                // When "To" is set, this will be the target of the button.
	Function        func(ctx app.Context) // If set, then the action will perform the function.
	Flat            bool                  // If true, then the button will be flat.
	BackgroundColor string                // The background color of the button.  If empty, then the default background color will be used.
	Color           string                // The color of the button.  If empty, then the default color will be used.
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

// performCancel performs the cancel function.
//
// If there is no cancel function, then nothing will be done.
func (c *blazarForm) performCancel(ctx app.Context) {
	// If there is no cancel function, then don't do anything.
	if c.ICancelFunction == nil {
		return
	}

	c.loading = true

	ctx.Async(func() {
		c.ICancelFunction(ctx)

		ctx.Dispatch(func(ctx app.Context) {
			c.loading = false
			ctx.Update()
		})
	})
}

// performSubmit performs the submit function.
//
// If there is no submit function, then then the *last* action will be done.
// If there is no last action, then nothing will be done.
func (c *blazarForm) performSubmit(ctx app.Context) {
	// If there is a submit function, then perform it.
	if c.ISubmitFunction != nil {
		c.loading = true

		ctx.Async(func() {
			c.ISubmitFunction(ctx)

			ctx.Dispatch(func(ctx app.Context) {
				c.loading = false
				ctx.Update()
			})
		})

		return
	}

	// If there are is at least one action, then perform the *last* one.
	if len(c.IActions) > 0 {
		lastAction := c.IActions[len(c.IActions)-1]

		c.performAction(ctx, lastAction)
	}
}

// performAction performs the action function.
func (c *blazarForm) performAction(ctx app.Context, action FormAction) {
	// If there is a function, then perform it.
	if action.Function != nil {
		c.loading = true

		ctx.Async(func() {
			action.Function(ctx)

			ctx.Dispatch(func(ctx app.Context) {
				c.loading = false
				ctx.Update()
			})
		})

		return
	}

	// If there is a link target, then navigate to it.
	if action.To != "" {
		ctx.Navigate(action.To)
	}
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

						if debugForm {
							slog.DebugContext(ctx.Context, "blazarForm: Keypress", "key", e.Get("key").String())
						}

						// If the user pressed "Enter", then perform the default action.
						//
						// If set, the default action is the submit function.
						// Otherwise, the default action is the *last* custom action.
						switch e.Get("key").String() {
						case "Enter":
							c.performSubmit(ctx)
						case "Escape":
							c.performCancel(ctx)
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
							Disabled(c.loading).
							Label(func() string {
								if c.ICancelLabel != "" {
									return c.ICancelLabel
								}
								return "Cancel"
							}()).
							Icon(c.ICancelIcon).
							On("click", func(ctx app.Context, e app.Event) {
								c.performCancel(ctx)
							})
					}),
					app.If(c.ISpacer, func() app.UI {
						return app.Span().Style("flex", "1")
					}),
					app.Range(c.IActions).Slice(func(i int) app.UI {
						action := c.IActions[i]
						button := Button().
							Flat(action.Flat).
							Disabled(c.loading).
							Label(action.Name).
							Icon(action.Icon).
							To(action.To).
							Target(action.Target).
							On("click", func(ctx app.Context, e app.Event) {
								c.performAction(ctx, action)
							})
						if action.BackgroundColor != "" {
							button = button.Style("background-color", action.BackgroundColor)
						}
						if action.Color != "" {
							button = button.Style("color", action.Color)
						}
						return button
					}),
					app.If(c.ISubmitFunction != nil, func() app.UI {
						return Button().
							Flat(false).
							Disabled(c.loading).
							Label(func() string {
								if c.ISubmitLabel != "" {
									return c.ISubmitLabel
								}
								return "Submit"
							}()).
							Icon(c.ISubmitIcon).
							On("click", func(ctx app.Context, e app.Event) {
								c.performSubmit(ctx)
							})
					}),
				),
		)
	for name, value := range c.IStyles {
		element = element.Style(name, value)
	}
	return element
}
