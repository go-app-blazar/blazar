package demo

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type AppBarPage struct {
	app.Compo
}

var _ app.Initializer = (*AppBarPage)(nil)

func (c *AppBarPage) OnInit() {
	slog.DebugContext(context.TODO(), "AppBarPage: OnInit")
}

func (c *AppBarPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "AppBarPage: OnNav")
}

func (c *AppBarPage) Render() app.UI {
	return blazar.Page().
		Body(
			blazar.AppBar(),
			blazar.AppBar().
				NoIcon(true).
				HeadlineText("App Bar with no icon"),
			blazar.AppBar().
				Icon(tabler.IconFilledMenu2).
				IconFunction(func(ctx app.Context) {
					app.Window().Call("alert", "Icon clicked")
				}).
				HeadlineText("App Bar with default icon"),
			blazar.AppBar().
				Icon(tabler.IconFilledMenu2).
				IconFunction(func(ctx app.Context) {
					app.Window().Call("alert", "Icon clicked")
				}).
				HeadlineText("App Bar with default icon and trailer").
				TrailerBody(
					app.Div().Text("Trailer"),
				),
			blazar.AppBar().
				Icon(tabler.IconFilledMenu2).
				IconFunction(func(ctx app.Context) {
					app.Window().Call("alert", "Icon clicked")
				}).
				HeadlineText("App Bar with default icon, subtitle, and trailer").
				SubtitleText("Subtitle text").
				TrailerBody(
					app.Div().Text("Trailer"),
				),
		)
}
