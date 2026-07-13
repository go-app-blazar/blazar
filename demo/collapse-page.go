package demo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type CollapsePage struct {
	app.Compo

	defaultOpen   bool
	defaultClosed bool
}

var _ app.Initializer = (*CollapsePage)(nil)

func (c *CollapsePage) OnInit() {
	slog.DebugContext(context.TODO(), "CollapsePage: OnInit")

	c.defaultOpen = true
	c.defaultClosed = false
}

func (c *CollapsePage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "CollapsePage: OnNav")
}

func (c *CollapsePage) Render() app.UI {
	return blazar.Page().
		Body(
			app.FieldSet().
				Style("display", "flex").
				Style("flex-direction", "column").
				Style("gap", "1em").
				Body(
					app.Legend().Text("Unbound"),
					blazar.Collapse().
						Open(true).
						Label("Default Open").
						SummaryText("This is open by default.").
						Body(
							app.Div().Text("This is the body of the collapse."),
						).
						OnOpenChange(func(ctx app.Context, open bool) {
							slog.InfoContext(ctx.Context, "CollapsePage: OnOpenChange", "open", open)
						}),
					blazar.Collapse().
						Open(false).
						Label("Default Closed").
						SummaryText("This is closed by default.").
						Body(
							app.Div().Text("This is the body of the collapse."),
						).
						OnOpenChange(func(ctx app.Context, open bool) {
							slog.InfoContext(ctx.Context, "CollapsePage: OnOpenChange", "open", open)
						}),
				),
			app.FieldSet().
				Style("display", "flex").
				Style("flex-direction", "column").
				Style("gap", "1em").
				Body(
					app.Legend().Text("Bound"),
					blazar.Collapse().
						Bind(&c.defaultOpen).
						Label("Default Open").
						SummaryText("Current state: "+fmt.Sprintf("%t", c.defaultOpen)).
						Body(
							app.Div().Text("This is the body of the collapse."),
						).
						OnOpenChange(func(ctx app.Context, open bool) {
							slog.InfoContext(ctx.Context, "CollapsePage: OnOpenChange", "open", open)
						}),
					blazar.Collapse().
						Bind(&c.defaultClosed).
						Label("Default Closed").
						SummaryText("Current state: "+fmt.Sprintf("%t", c.defaultClosed)).
						Body(
							app.Div().Text("This is the body of the collapse."),
						).
						OnOpenChange(func(ctx app.Context, open bool) {
							slog.InfoContext(ctx.Context, "CollapsePage: OnOpenChange", "open", open)
						}),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Bound Values"),
					app.Div().Text(fmt.Sprintf("Default Open: %t", c.defaultOpen)),
					app.Div().Text(fmt.Sprintf("Default Closed: %t", c.defaultClosed)),
				),
		)
}
