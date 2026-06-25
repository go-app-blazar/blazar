package blazar

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Collapse() *blazarCollapse {
	return &blazarCollapse{}
}

type blazarCollapse struct {
	app.Compo
	UseEvents
	ILabel       string
	IDisabled    bool
	ISummaryText string
	ISummary     []app.UI
	IBody        []app.UI
	BindOpen     *bool
}

var _ app.Composer = (*blazarCollapse)(nil)
var _ app.Updater = (*blazarCollapse)(nil)

func (c *blazarCollapse) OnUpdate(ctx app.Context) {
	if debugCollapse {
		slog.DebugContext(ctx.Context, "blazarCollapse: OnUpdate")
	}
	if c.BindOpen != nil {
		if debugCollapse {
			slog.DebugContext(ctx.Context, "blazarCollapse: OnUpdate", "*BindValue", *c.BindOpen)
		}
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
	if c.BindOpen == nil {
		c.BindOpen = new(bool)
	}
	*c.BindOpen = open
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

func (c *blazarCollapse) On(event string, function func(ctx app.Context, e app.Event)) *blazarCollapse {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarCollapse) Bind(variable *bool) *blazarCollapse {
	c.BindOpen = variable
	return c
}

func (c *blazarCollapse) Render() app.UI {
	open := false
	if c.BindOpen != nil {
		open = *c.BindOpen
	}
	if debugCollapse {
		slog.DebugContext(context.TODO(), "blazarCollapse: Render", "BindOpen", c.BindOpen, "open", open)
	}

	var element app.UI

	disabledClass := ""
	if c.IDisabled {
		disabledClass = "disabled"
	}

	closedIcon := "chevron-down"
	closedClass := ""
	if !open {
		closedIcon = "chevron-right"
		closedClass = "closed"
	}

	element = app.Div().
		Class("blazar-collapse").
		Class(disabledClass).
		Class(closedClass).
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
						Class("blazar-collapse-icon").
						Body(
							Icon().
								Icon(closedIcon),
						),
				).
				On("click", func(ctx app.Context, e app.Event) {
					open = !open
					if c.BindOpen != nil {
						*c.BindOpen = open
					}
					if debugCollapse {
						slog.DebugContext(ctx.Context, "Collapse: OnClick", "BindOpen", c.BindOpen, "open", open)
					}
					ctx.Update()
				}),
			app.If(open, func() app.UI {
				return app.Div().
					Class("blazar-collapse-content").
					Body(c.IBody...)
			}),
		)
	if !c.IDisabled {
		element = c.UseEvents.Wrap(element)
	}
	return element
}
