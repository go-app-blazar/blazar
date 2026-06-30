package blazar

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type blazarAppBar struct {
	app.Compo

	IClasses      []string
	INoIcon       bool
	IIcon         string
	IIconFunction func(ctx app.Context, e app.Event)

	IHeadline         string
	IHeadlineFunction func() app.UI
	ISubtitle         string
	ISubtitleFunction func() app.UI
	ITrailerFunction  func() app.UI
}

func AppBar() *blazarAppBar {
	return &blazarAppBar{
		IIcon: "bars",
	}
}

func (c *blazarAppBar) Class(class ...string) *blazarAppBar {
	c.IClasses = class
	return c
}

func (c *blazarAppBar) NoIcon(noIcon bool) *blazarAppBar {
	c.INoIcon = noIcon
	return c
}

func (c *blazarAppBar) Icon(icon string) *blazarAppBar {
	c.IIcon = icon
	return c
}

func (c *blazarAppBar) IconFunction(function func(ctx app.Context, e app.Event)) *blazarAppBar {
	c.IIconFunction = function
	return c
}

func (c *blazarAppBar) HeadlineText(text string) *blazarAppBar {
	c.IHeadline = text
	return c
}

func (c *blazarAppBar) HeadlineFunction(function func() app.UI) *blazarAppBar {
	c.IHeadlineFunction = function
	return c
}

func (c *blazarAppBar) SubtitleText(text string) *blazarAppBar {
	c.ISubtitle = text
	return c
}

func (c *blazarAppBar) SubtitleFunction(function func() app.UI) *blazarAppBar {
	c.ISubtitleFunction = function
	return c
}

func (c *blazarAppBar) TrailerFunction(function func() app.UI) *blazarAppBar {
	c.ITrailerFunction = function
	return c
}

func (c *blazarAppBar) Render() app.UI {
	if debugAppBar {
		slog.DebugContext(context.TODO(), "AppBar: Render", "INoIcon", c.INoIcon)
	}

	return app.Div().
		Class(append([]string{"blazar-app-bar"}, c.IClasses...)...).
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
									c.IIconFunction(ctx, e)
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
							app.If(c.IHeadlineFunction != nil, func() app.UI {
								return c.IHeadlineFunction()
							}).Else(func() app.UI {
								return app.Div().
									Class("blazar-app-bar__headline-text").
									Text(c.IHeadline)
							}),
						),
					app.If(c.ISubtitle != "" || c.ISubtitleFunction != nil, func() app.UI {
						return app.Div().
							Class("blazar-app-bar__subtitle").
							Body(
								app.If(c.ISubtitleFunction != nil, func() app.UI {
									return c.ISubtitleFunction()
								}).Else(func() app.UI {
									return app.Div().
										Class("blazar-app-bar__subtitle-text").
										Text(c.ISubtitle)
								}),
							)
					}),
				),
			app.If(c.ITrailerFunction != nil, func() app.UI {
				return app.Div().
					Class("blazar-app-bar__trailer").
					Body(c.ITrailerFunction())
			}),
		)
}
