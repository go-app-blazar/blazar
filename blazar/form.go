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

	ISpacer  bool
	IBody    []app.UI
	IActions []FormAction

	loading bool
}

type FormAction struct {
	Name     string                // The name of the action.
	Icon     string                // The icon of the action.  If empty, then no icon will be shown.
	To       string                // If set, then the action will navigate to the target URL.
	Target   string                // When "To" is set, this will be the target of the button.
	Function func(ctx app.Context) // If set, then the action will perform the function.
	Flat     bool                  // If true, then the button will be flat.
	Outline  bool                  // If true, then the button will be outlined.
	Color    string                // The color of the button.  If empty, then the primary theme color will be used.
	Submit   bool                  // If true, then the action will be the submit button.
	Cancel   bool                  // If true, then the action will be the cancel button.
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
	var cancelAction *FormAction

	// If there is a cancel action, then use the last one given.
	for _, action := range c.IActions {
		if action.Cancel {
			cancelAction = &action
			// Keep going; we'll keep the last one.
		}
	}
	// If there is no cancel action, then nothing will be done.
	if cancelAction == nil {
		return
	}

	c.performAction(ctx, *cancelAction)
}

// performSubmit performs the submit function.
//
// If there is no submit function, then then the *last* action will be done.
// If there is no last action, then nothing will be done.
func (c *blazarForm) performSubmit(ctx app.Context) {
	var submitAction *FormAction

	// If there is a submit action, then use the last one given.
	for _, action := range c.IActions {
		if action.Submit {
			submitAction = &action
			// Keep going; we'll keep the last one.
		}
	}
	// If there is no submit action, then use the last normal action.
	if submitAction == nil {
		for _, action := range c.IActions {
			if !action.Cancel {
				submitAction = &action
				// Keep going; we'll keep the last one.
			}
		}
	}

	// If there is no submit action, then nothing will be done.
	if submitAction == nil {
		return
	}

	c.performAction(ctx, *submitAction)
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

	// TODO: Handle a "To" action.
}

func (c *blazarForm) Render() app.UI {
	formDisplay := "block"
	if len(c.IBody) == 0 {
		formDisplay = "none"
	}

	var cancelAction *FormAction
	var submitAction *FormAction
	var otherActions []FormAction
	for _, action := range c.IActions {
		if action.Cancel {
			cancelAction = &action
			// Keep going; we'll keep the last one.
			continue
		}
		if action.Submit {
			submitAction = &action
			// Keep going; we'll keep the last one.
			continue
		}
		otherActions = append(otherActions, action)
	}
	if submitAction == nil {
		if len(otherActions) > 0 {
			submitAction = &otherActions[len(otherActions)-1]
			otherActions = otherActions[:len(otherActions)-1]
		}
	}

	element := app.Div().
		Class(append([]string{"blazar-form"}, c.IClasses...)...).
		Body(
			c.UseEvents.Wrap(
				app.Div().
					Class("blazar-form__form").
					Style("display", formDisplay).
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
							ctx.Update()
						case "Escape":
							c.performCancel(ctx)
							ctx.Update()
						}
					}).
					Body(
						c.IBody...,
					),
			),
			app.Div().
				Class("blazar-form__actions").
				Body(
					app.If(cancelAction != nil, func() app.UI {
						action := cancelAction

						button := Button().
							Flat(true).
							Disabled(c.loading).
							Label(func() string {
								if action.Name != "" {
									return action.Name
								}
								return "Cancel"
							}()).
							Icon(action.Icon).
							On("click", func(ctx app.Context, e app.Event) {
								c.performCancel(ctx)
							})
						if action.Color == "" {
							action.Color = "var(--blazar-theme-primary)"
						}
						if action.Flat {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						} else if action.Outline {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						} else {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						}
						return button
					}),
					app.If(c.ISpacer, func() app.UI {
						return app.Span().Style("flex", "1")
					}),
					app.Range(otherActions).Slice(func(i int) app.UI {
						action := otherActions[i]
						button := Button().
							Flat(action.Flat).
							Outline(action.Outline).
							Disabled(c.loading).
							Label(action.Name).
							Icon(action.Icon).
							To(action.To).
							Target(action.Target).
							On("click", func(ctx app.Context, e app.Event) {
								c.performAction(ctx, action)
							})
						if action.Color == "" {
							action.Color = "var(--blazar-theme-secondary)"
						}
						if action.Flat {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						} else if action.Outline {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						} else {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						}
						return button
					}),
					app.If(submitAction != nil, func() app.UI {
						action := submitAction

						button := Button().
							Flat(false).
							Disabled(c.loading).
							Label(func() string {
								if action.Name != "" {
									return action.Name
								}
								return "Submit"
							}()).
							Icon(action.Icon).
							On("click", func(ctx app.Context, e app.Event) {
								c.performSubmit(ctx)
							})
						if action.Color == "" {
							action.Color = "var(--blazar-theme-primary)"
						}
						if action.Flat {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						} else if action.Outline {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						} else {
							if action.Color != "" {
								button = button.Style("--color", action.Color)
							}
						}
						return button
					}),
				),
		)
	for name, value := range c.IStyles {
		element = element.Style(name, value)
	}
	return element
}
