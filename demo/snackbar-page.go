package demo

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type SnackbarPage struct {
	app.Compo

	paragraphs uint
}

var _ app.Initializer = (*SnackbarPage)(nil)

func (c *SnackbarPage) OnInit() {
	slog.DebugContext(context.TODO(), "SnackbarPage: OnInit")

	c.paragraphs = 2
}

func (c *SnackbarPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "SnackbarPage: OnNav")
}

func (c *SnackbarPage) Render() app.UI {
	var paragraphs []string
	for range c.paragraphs {
		paragraphs = append(paragraphs, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	}

	return blazar.Page().
		Body(
			blazar.Snackbar().
				Observe("snackbar-message"),
			app.FieldSet().
				Body(
					app.Legend().Text("Configuration"),
					app.Div().Body(
						blazar.Button().
							Label("Show Snackbar").
							OnClick(func(ctx app.Context, e htmlevent.MouseEvent) {
								blazar.SendSnackbarMessage(ctx, "snackbar-message", blazar.SnackbarMessage().Text("Hello, world!"))
							}),
						blazar.Button().
							Label("Show Bad Snackbar").
							OnClick(func(ctx app.Context, e htmlevent.MouseEvent) {
								blazar.SendSnackbarMessage(ctx, "snackbar-message", blazar.SnackbarMessage().Text("Oh no, world!").Bad())
							}),
						blazar.Input[uint]().
							Label("Paragraphs").
							Bind(&c.paragraphs),
					),
				),
			app.Range(paragraphs).Slice(func(i int) app.UI {
				paragraph := paragraphs[i]

				return app.P().
					Text(paragraph)
			}),
			blazar.Button().
				Label("Show Snackbar").
				OnClick(func(ctx app.Context, e htmlevent.MouseEvent) {
					blazar.SendSnackbarMessage(ctx, "snackbar-message", blazar.SnackbarMessage().Text("Goodbye, world!"))
				}),
		)
}
