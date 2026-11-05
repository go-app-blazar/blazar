package demo

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type SnackbarPage struct {
	app.Compo

	paragraphs        uint
	message           string
	durationInSeconds uint
	closable          bool
	withAction        bool
}

var _ app.Initializer = (*SnackbarPage)(nil)

func (c *SnackbarPage) OnInit() {
	slog.DebugContext(context.TODO(), "SnackbarPage: OnInit")

	c.paragraphs = 2
	c.message = "Hello, World!"
	c.durationInSeconds = 10
	c.closable = false
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
						blazar.Input[uint]().
							Label("Page Paragraphs").
							Bind(&c.paragraphs),
						blazar.Input[string]().
							Label("Snackbar Message").
							Bind(&c.message),
						app.Div().
							Style("font-size", "0.7em").
							Style("display", "flex").
							Style("flex-direction", "row").
							Style("gap", "1em").
							Body(
								app.A().
									Text("Small").
									OnClick(htmlevent.WrapMouseEvent(func(ctx app.Context, e htmlevent.MouseEvent) {
										c.message = "Hello, World!"
									})),
								app.A().
									Text("Medium").
									OnClick(htmlevent.WrapMouseEvent(func(ctx app.Context, e htmlevent.MouseEvent) {
										c.message = "This is a medium-sized message."
									})),
								app.A().
									Text("Large").
									OnClick(htmlevent.WrapMouseEvent(func(ctx app.Context, e htmlevent.MouseEvent) {
										c.message = "This is a large-sized message. It is used to display important information that needs to be read by the user."
									})),
							),
						blazar.Input[uint]().
							Label("Snackbar Duration").
							Bind(&c.durationInSeconds),
						blazar.Input[bool]().
							Label("Closable").
							Bind(&c.closable),
						blazar.Input[bool]().
							Label("With Action").
							Bind(&c.withAction),
					),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Actions"),
					app.Div().Body(
						blazar.Button().
							Label("Show Snackbar").
							OnClick(func(ctx app.Context, e htmlevent.MouseEvent) {
								message := blazar.SnackbarMessage().
									Text(c.message).
									Duration(time.Duration(c.durationInSeconds) * time.Second).
									Closable(c.closable)
								if c.withAction {
									message.Action(blazar.SnackbarAction{
										Name: "Action",
										Icon: tabler.IconFilledCheck,
										Function: func(ctx app.Context) {
											slog.DebugContext(ctx, "Action clicked")
										},
									})
								}

								blazar.SendSnackbarMessage(ctx, "snackbar-message", message)
							}),
						blazar.Button().
							Label("Show Good Snackbar").
							OnClick(func(ctx app.Context, e htmlevent.MouseEvent) {
								message := blazar.SnackbarMessage().
									Text(c.message).
									Duration(time.Duration(c.durationInSeconds) * time.Second).
									Closable(c.closable)
								if c.withAction {
									message.Action(blazar.SnackbarAction{
										Name: "Action",
										Icon: tabler.IconFilledCheck,
										Function: func(ctx app.Context) {
											slog.DebugContext(ctx, "Action clicked")
										},
									})
								}

								message.Good()

								blazar.SendSnackbarMessage(ctx, "snackbar-message", message)
							}),
						blazar.Button().
							Label("Show Bad Snackbar").
							OnClick(func(ctx app.Context, e htmlevent.MouseEvent) {
								message := blazar.SnackbarMessage().
									Text(c.message).
									Duration(time.Duration(c.durationInSeconds) * time.Second).
									Closable(c.closable)
								if c.withAction {
									message.Action(blazar.SnackbarAction{
										Name: "Action",
										Icon: tabler.IconFilledCheck,
										Function: func(ctx app.Context) {
											slog.DebugContext(ctx, "Action clicked")
										},
									})
								}

								message.Bad()

								blazar.SendSnackbarMessage(ctx, "snackbar-message", message)
							}),
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
					message := blazar.SnackbarMessage().
						Text(c.message).
						Duration(time.Duration(c.durationInSeconds) * time.Second).
						Closable(c.closable)
					if c.withAction {
						message.Action(blazar.SnackbarAction{
							Name: "Action",
							Icon: tabler.IconFilledCheck,
							Function: func(ctx app.Context) {
								slog.DebugContext(ctx, "Action clicked")
							},
						})
					}

					blazar.SendSnackbarMessage(ctx, "snackbar-message", message)
				}),
		)
}
