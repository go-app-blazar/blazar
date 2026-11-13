package blazar

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type blazarAppBar struct {
	app.Compo
	MixHTML

	INoIcon       bool
	IIcon         string
	IIconFunction func(ctx app.Context)

	IHeadlineText string
	IHeadlineBody []app.UI
	ISubtitleText string
	ISubtitleBody []app.UI
	ITrailerBody  []app.UI
}

func AppBar() *blazarAppBar {
	return &blazarAppBar{
		IIcon: tabler.IconFilledMenu2,
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
	c.IHeadlineText = text
	return c
}

func (c *blazarAppBar) HeadlineBody(body ...app.UI) *blazarAppBar {
	c.IHeadlineBody = body
	return c
}

func (c *blazarAppBar) SubtitleText(text string) *blazarAppBar {
	c.ISubtitleText = text
	return c
}

func (c *blazarAppBar) SubtitleBody(body ...app.UI) *blazarAppBar {
	c.ISubtitleBody = body
	return c
}

func (c *blazarAppBar) TrailerBody(body ...app.UI) *blazarAppBar {
	c.ITrailerBody = body
	return c
}

func (c *blazarAppBar) Render() app.UI {
	if debugAppBar {
		slog.DebugContext(context.TODO(), "AppBar: Render", "INoIcon", c.INoIcon)
	}

	var finalElement app.UI
	finalElement = app.Div().
		Class("blazar-app-bar").
		Style("width", "100%").
		Body(
			app.If(!c.INoIcon && c.IIcon != "" && c.IIconFunction != nil, func() app.UI {
				return app.Div().
					Class("blazar-app-bar__icon").
					Body(
						Icon().
							Icon(c.IIcon).
							OnClick(func(ctx app.Context, e htmlevent.MouseEvent) {
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
							app.IfSlice(len(c.IHeadlineBody) > 0, func() []app.UI {
								return c.IHeadlineBody
							}).Else(func() app.UI {
								return app.Div().
									Class("blazar-app-bar__headline-text").
									Text(c.IHeadlineText)
							}),
						),
					app.If(c.ISubtitleText != "" || len(c.ISubtitleBody) > 0, func() app.UI {
						return app.Div().
							Class("blazar-app-bar__subtitle").
							Body(
								app.IfSlice(len(c.ISubtitleBody) > 0, func() []app.UI {
									return c.ISubtitleBody
								}).Else(func() app.UI {
									return app.Div().
										Class("blazar-app-bar__subtitle-text").
										Text(c.ISubtitleText)
								}),
							)
					}),
				),
			app.If(len(c.ITrailerBody) > 0, func() app.UI {
				return app.Div().
					Class("blazar-app-bar__trailer").
					Body(c.ITrailerBody...)
			}),
		)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
