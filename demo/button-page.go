package demo

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type ButtonPage struct {
	app.Compo
}

var _ app.Initializer = (*ButtonPage)(nil)

func (c *ButtonPage) OnInit() {
	slog.DebugContext(context.TODO(), "ButtonPage: OnInit")
}

func (c *ButtonPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "ButtonPage: OnNav")
}

func (c *ButtonPage) Render() app.UI {
	clickFunction := func(ctx app.Context, e htmlevent.MouseEvent) {
		slog.DebugContext(ctx.Context, "ButtonPage: Click", "event", e)
		app.Window().Call("alert", "Button clicked")
	}

	return blazar.Page().
		Body(
			app.FieldSet().
				Body(
					app.Legend().Text("Navigation"),
					app.Div().Body(
						blazar.Button().
							Label("Default").
							To("/"),
						blazar.Button().
							Label("Flat").
							Flat(true).
							To("/"),
						blazar.Button().
							Label("Outline").
							Outline(true).
							Flat(true).
							To("/"),
						blazar.Button().
							Label("Round").
							Round(true).
							To("/"),
						blazar.Button().
							Label("Disabled").
							Disabled(true).
							To("/"),
					),
					app.Div().Body(
						blazar.Button().
							Label("Default").
							Icon(tabler.IconFilledHome).
							To("/"),
						blazar.Button().
							Label("Flat").
							Flat(true).
							Icon(tabler.IconFilledHome).
							To("/"),
						blazar.Button().
							Label("Outline").
							Outline(true).
							Flat(true).
							Icon(tabler.IconFilledHome).
							To("/"),
						blazar.Button().
							Label("Round").
							Round(true).
							Icon(tabler.IconFilledHome).
							To("/"),
						blazar.Button().
							Label("Disabled").
							Disabled(true).
							Icon(tabler.IconFilledHome).
							To("/"),
					),
					app.Div().Body(
						blazar.Button().
							Icon(tabler.IconFilledHome).
							To("/"),
						blazar.Button().
							Flat(true).
							Icon(tabler.IconFilledHome).
							To("/"),
						blazar.Button().
							Outline(true).
							Icon(tabler.IconFilledHome).
							To("/"),
						blazar.Button().
							Round(true).
							Icon(tabler.IconFilledHome).
							To("/"),
						blazar.Button().
							Disabled(true).
							Icon(tabler.IconFilledHome).
							To("/"),
					),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Action"),
					app.Div().Body(
						blazar.Button().
							Label("Default").
							OnClick(clickFunction),
						blazar.Button().
							Label("Flat").
							Flat(true).
							OnClick(clickFunction),
						blazar.Button().
							Label("Outline").
							Outline(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Round").
							Round(true).
							OnClick(clickFunction),
						blazar.Button().
							Label("Disabled").
							Disabled(true).
							OnClick(clickFunction),
					),
					app.Div().Body(
						blazar.Button().
							Label("Default").
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Flat").
							Flat(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Outline").
							Outline(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Round").
							Round(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Disabled").
							Disabled(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
					),
					app.Div().Body(
						blazar.Button().
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Flat(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Outline(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Round(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
						blazar.Button().
							Disabled(true).
							Icon(tabler.IconFilledHome).
							OnClick(clickFunction),
					),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Auto Focus"),
					app.Div().Body(
						blazar.Button().
							Label("Auto Focus").
							AutoFocus(true).
							OnClick(clickFunction),
					),
				),
		)
}
