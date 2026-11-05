package demo

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type StatusBarPage struct {
	app.Compo
}

var _ app.Initializer = (*StatusBarPage)(nil)

func (c *StatusBarPage) OnInit() {
	slog.DebugContext(context.TODO(), "StatusBarPage: OnInit")
}

func (c *StatusBarPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "StatusBarPage: OnNav")
}

func (c *StatusBarPage) Render() app.UI {
	return blazar.Page().
		Body(
			blazar.StatusBar(),
			blazar.StatusBar().
				Neutral().
				Text("Neutral StatusBar body text"),
			blazar.StatusBar().
				Neutral().
				Text("Neutral StatusBar body text").
				Action(
					blazar.StatusBarAction{
						Name: "Action 1",
						Icon: tabler.IconOutlineInfoCircle,
						Function: func(ctx app.Context) {
							slog.InfoContext(ctx.Context, "StatusBarPage: Action 1")
						},
					},
					blazar.StatusBarAction{
						Name: "",
						Icon: tabler.IconFilledX,
						Function: func(ctx app.Context) {
							slog.InfoContext(ctx.Context, "StatusBarPage: Action 2")
						},
					},
				),
			blazar.StatusBar().
				Good().
				Text("Good StatusBar body text"),
			blazar.StatusBar().
				Good().
				Text("Good StatusBar body text").
				Action(
					blazar.StatusBarAction{
						Name: "Action 1",
						Icon: tabler.IconOutlineInfoCircle,
						Function: func(ctx app.Context) {
							slog.InfoContext(ctx.Context, "StatusBarPage: Action 1")
						},
					},
					blazar.StatusBarAction{
						Name: "",
						Icon: tabler.IconFilledX,
						Function: func(ctx app.Context) {
							slog.InfoContext(ctx.Context, "StatusBarPage: Action 2")
						},
					},
				),
			blazar.StatusBar().
				Bad().
				Text("Bad StatusBar body text"),
			blazar.StatusBar().
				Bad().
				Text("Bad StatusBar body text").
				Action(
					blazar.StatusBarAction{
						Name: "Action 1",
						Icon: tabler.IconOutlineInfoCircle,
						Function: func(ctx app.Context) {
							slog.InfoContext(ctx.Context, "StatusBarPage: Action 1")
						},
					},
					blazar.StatusBarAction{
						Name: "",
						Icon: tabler.IconFilledX,
						Function: func(ctx app.Context) {
							slog.InfoContext(ctx.Context, "StatusBarPage: Action 2")
						},
					},
				),
		)
}
