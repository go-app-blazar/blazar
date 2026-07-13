package blazar

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/deref"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Collapse() *blazarCollapse {
	return &blazarCollapse{}
}

type blazarCollapse struct {
	app.Compo
	ILabel        string
	IDisabled     bool
	ISummaryText  string
	ISummary      []app.UI
	IBody         []app.UI
	IOnOpenChange func(ctx app.Context, open bool)
	open          bool
	bindOpen      *bool
}

var _ app.Composer = (*blazarCollapse)(nil)
var _ app.Updater = (*blazarCollapse)(nil)

func (c *blazarCollapse) OnUpdate(ctx app.Context) {
	if debugCollapse {
		slog.DebugContext(ctx.Context, "blazarCollapse: OnUpdate")
	}
	if c.bindOpen != nil {
		if debugCollapse {
			slog.DebugContext(ctx.Context, "blazarCollapse: OnUpdate", "*BindValue", *c.bindOpen)
		}
		c.open = *c.bindOpen
	} else {
		if debugCollapse {
			slog.DebugContext(ctx.Context, "blazarCollapse: OnUpdate: BindOpen is nil.")
		}
	}
}

func (c *blazarCollapse) Disabled(disabled bool) *blazarCollapse {
	c.IDisabled = disabled
	return c
}

func (c *blazarCollapse) Open(open bool) *blazarCollapse {
	if debugCollapse {
		slog.DebugContext(context.TODO(), "blazarCollapse: Open", "open", open)
	}
	c.open = open
	if c.bindOpen != nil {
		*c.bindOpen = open
	}
	return c
}

func (c *blazarCollapse) Label(label string) *blazarCollapse {
	c.ILabel = label
	return c
}

func (c *blazarCollapse) SummaryText(summaryText string) *blazarCollapse {
	c.ISummaryText = summaryText
	return c
}

func (c *blazarCollapse) Summary(summary ...app.UI) *blazarCollapse {
	c.ISummary = summary
	return c
}

func (c *blazarCollapse) Body(body ...app.UI) *blazarCollapse {
	c.IBody = body
	return c
}

// OnOpenChange is called when the collapse is opened or closed.
// The new value is passed to the function.
func (c *blazarCollapse) OnOpenChange(function func(ctx app.Context, open bool)) *blazarCollapse {
	c.IOnOpenChange = function
	return c
}

func (c *blazarCollapse) Bind(variable *bool) *blazarCollapse {
	if debugCollapse {
		slog.DebugContext(context.TODO(), "blazarCollapse: Bind", "variable", deref.String(variable))
	}
	c.bindOpen = variable
	if c.bindOpen != nil {
		c.open = *c.bindOpen
	}
	return c
}

func (c *blazarCollapse) Render() app.UI {
	if debugCollapse {
		slog.DebugContext(context.TODO(), "blazarCollapse: Render", "bindOpen", deref.String(c.bindOpen), "open", c.open)
	}

	var element app.UI

	disabledClass := ""
	if c.IDisabled {
		disabledClass = "disabled"
	}

	closedIcon := "chevron-down"
	closedClass := ""
	if !c.open {
		closedIcon = "chevron-right"
		closedClass = "closed"
	}

	element = app.Div().
		Class("blazar-collapse", disabledClass, closedClass).
		Body(
			app.Div().
				Class("blazar-collapse__top").
				Style("cursor", "pointer").
				Body(
					app.Span().
						Class("blazar-collapse__label").
						Text(c.ILabel),
					app.If(len(c.ISummary) > 0, func() app.UI {
						return app.Span().
							Class("blazar-collapse__summary").
							Body(c.ISummary...)
					}).Else(func() app.UI {
						return app.Span().
							Class("blazar-collapse__summary-text").
							Text(c.ISummaryText).
							Title(c.ISummaryText)
					}),
					app.Span().Style("flex", "1"),
					app.Span().
						Class("blazar-collapse__icon").
						Body(
							Icon().
								Icon(closedIcon),
						),
				).
				On("click", func(ctx app.Context, e app.Event) {
					c.open = !c.open
					if c.bindOpen != nil {
						*c.bindOpen = c.open
					}
					if debugCollapse {
						slog.DebugContext(ctx.Context, "Collapse: OnClick", "BindOpen", deref.String(c.bindOpen), "open", c.open)
					}
					if c.IOnOpenChange != nil {
						c.IOnOpenChange(ctx, c.open)
					}
					ctx.Update()
				}),
			app.If(c.open, func() app.UI {
				return app.Div().
					Class("blazar-collapse-content").
					Body(c.IBody...)
			}),
		)
	return element
}
