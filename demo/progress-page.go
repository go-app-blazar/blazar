package demo

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type ProgressPage struct {
	app.Compo

	animated      bool
	indeterminate bool
	percentage    uint
}

func (c *ProgressPage) OnMount(ctx app.Context) {
	slog.DebugContext(ctx.Context, "ProgressPage: OnMount")

	c.animated = false
	c.indeterminate = false
	c.percentage = 60
}

func (c *ProgressPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "ProgressPage: OnNav")
}

func (c *ProgressPage) OnUpdate(ctx app.Context) {
	slog.DebugContext(ctx.Context, "ProgressPage: OnUpdate", "percentage", c.percentage, "indeterminate", c.indeterminate)
}

func (c *ProgressPage) Render() app.UI {
	slog.DebugContext(context.TODO(), "ProgressPage: Render", "percentage", c.percentage, "indeterminate", c.indeterminate)

	return blazar.Page().
		Body(
			app.FieldSet().
				Body(
					app.Legend().Text("Configuration"),
					blazar.Input[uint]().
						Label("Percentage").
						Min(0).
						Max(100).
						Bind(&c.percentage),
					blazar.Input[bool]().
						Label("Animated").
						Bind(&c.animated),
					blazar.Input[bool]().
						Label("Indeterminate").
						Bind(&c.indeterminate),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Progress Bars"),
					app.Div().
						Style("margin-bottom", "1em").
						Body(
							app.Div().Text("Default"),
							blazar.ProgressBar().
								Percentage(c.percentage).
								Animated(c.animated).
								Indeterminate(c.indeterminate),
						),
					app.Div().
						Style("margin-bottom", "1em").
						Body(
							app.Div().Text("Height: 1px"),
							blazar.ProgressBar().
								Height("1px").
								Percentage(c.percentage).
								Animated(c.animated).
								Indeterminate(c.indeterminate),
						),
					app.Div().
						Style("margin-bottom", "1em").
						Body(
							app.Div().Text("Height: 1em"),
							blazar.ProgressBar().
								Height("1em").
								Percentage(c.percentage).
								Animated(c.animated).
								Indeterminate(c.indeterminate),
						),
				),
		)
}
