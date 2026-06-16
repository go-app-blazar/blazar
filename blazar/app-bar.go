package blazar

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type blazarAppBar struct {
	app.Compo

	INoIcon       bool
	IIcon         string
	IIconFunction func(ctx app.Context)

	IHeadline   string
	IHeadlineUI app.UI
	ISubtitle   string
	ISubtitleUI app.UI
	ITrailer    app.UI
}

func AppBar() *blazarAppBar {
	return &blazarAppBar{
		IIcon: "bars",
	}
}

func (c *blazarAppBar) NoIcon(noIcon bool) *blazarAppBar {
	c.INoIcon = noIcon
	return c
}

func (c *blazarAppBar) Icon(icon string) *blazarAppBar {
	c.IIcon = icon
	return c
}

func (c *blazarAppBar) IconFunction(function func(ctx app.Context)) *blazarAppBar {
	c.IIconFunction = function
	return c
}

func (c *blazarAppBar) HeadlineText(text string) *blazarAppBar {
	c.IHeadline = text
	return c
}

func (c *blazarAppBar) HeadlineUI(ui app.UI) *blazarAppBar {
	c.IHeadlineUI = ui
	return c
}

func (c *blazarAppBar) SubtitleText(text string) *blazarAppBar {
	c.ISubtitle = text
	return c
}

func (c *blazarAppBar) SubtitleUI(ui app.UI) *blazarAppBar {
	c.ISubtitleUI = ui
	return c
}

func (c *blazarAppBar) Trailer(ui app.UI) *blazarAppBar {
	c.ITrailer = ui
	return c
}

func (c *blazarAppBar) Render() app.UI {
	slog.InfoContext(context.TODO(), "AppBar: Render")

	return app.Div().
		Class("blazar-app-bar").
		Style("width", "100%").
		Body(
			app.If(!c.INoIcon && c.IIcon != "" && c.IIconFunction != nil, func() app.UI {
				return app.Div().
					Class("blazar-app-bar__icon").
					Body(
						Icon().
							Icon(c.IIcon).
							On("click", func(ctx app.Context, e app.Event) {
								if c.IIconFunction != nil {
									c.IIconFunction(ctx)
								}
							}),
					)
			}),
			app.Div().
				Class("blazar-app-bar__content").
				Body(
					app.Div().
						Class("blazar-app-bar__headline").
						Body(
							app.If(c.IHeadlineUI != nil, func() app.UI {
								return c.IHeadlineUI
							}).Else(func() app.UI {
								return app.Div().
									Class("blazar-app-bar__headline-text").
									Text(c.IHeadline)
							}),
						),
					app.If(c.ISubtitle != "" || c.ISubtitleUI != nil, func() app.UI {
						return app.Div().
							Class("blazar-app-bar__subtitle").
							Body(
								app.If(c.ISubtitleUI != nil, func() app.UI {
									return c.ISubtitleUI
								}).Else(func() app.UI {
									return app.Div().
										Class("blazar-app-bar__subtitle-text").
										Text(c.ISubtitle)
								}),
							)
					}),
				),
			app.If(c.ITrailer != nil, func() app.UI {
				return app.Div().
					Class("blazar-app-bar__trailer").
					Body(c.ITrailer)
			}),
		)
}
