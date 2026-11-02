package demo

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/go-app-blazar/heroicons-font"
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
							Icon(heroicons.IconSolidHome).
							To("/"),
						blazar.Button().
							Label("Flat").
							Flat(true).
							Icon(heroicons.IconSolidHome).
							To("/"),
						blazar.Button().
							Label("Outline").
							Outline(true).
							Flat(true).
							Icon(heroicons.IconSolidHome).
							To("/"),
						blazar.Button().
							Label("Round").
							Round(true).
							Icon(heroicons.IconSolidHome).
							To("/"),
						blazar.Button().
							Label("Disabled").
							Disabled(true).
							Icon(heroicons.IconSolidHome).
							To("/"),
					),
					app.Div().Body(
						blazar.Button().
							Icon(heroicons.IconSolidHome).
							To("/"),
						blazar.Button().
							Flat(true).
							Icon(heroicons.IconSolidHome).
							To("/"),
						blazar.Button().
							Outline(true).
							Icon(heroicons.IconSolidHome).
							To("/"),
						blazar.Button().
							Round(true).
							Icon(heroicons.IconSolidHome).
							To("/"),
						blazar.Button().
							Disabled(true).
							Icon(heroicons.IconSolidHome).
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
							Icon(heroicons.IconSolidHome).
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
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Flat").
							Flat(true).
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Outline").
							Outline(true).
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Round").
							Round(true).
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
						blazar.Button().
							Label("Disabled").
							Disabled(true).
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
					),
					app.Div().Body(
						blazar.Button().
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
						blazar.Button().
							Flat(true).
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
						blazar.Button().
							Outline(true).
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
						blazar.Button().
							Round(true).
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
						blazar.Button().
							Disabled(true).
							Icon(heroicons.IconSolidHome).
							OnClick(clickFunction),
					),
				),
		)
}
